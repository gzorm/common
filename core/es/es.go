package es

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type ElasticsearchClient struct {
	client *elasticsearch8.Client
}

// NewElasticsearchClient 初始化Elasticsearch客户端
func NewElasticsearchClient(useTLS bool, caCertPath, username, password string, addresses []string) (*ElasticsearchClient, error) {
	var transport *http.Transport
	if useTLS {
		// 读取本地证书文件
		cert, err := os.ReadFile(caCertPath)
		if err != nil {
			return nil, fmt.Errorf("Error reading CA certificate: %s", err)
		}

		// 创建一个新的证书池，并添加CA证书
		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(cert); !ok {
			return nil, fmt.Errorf("Failed to append CA certificate")
		}

		// 配置TLS客户端，使用CA证书池
		tlsConfig := &tls.Config{
			RootCAs: caCertPool,
		}

		transport = &http.Transport{
			TLSClientConfig:       tlsConfig,
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
		}
	} else {
		transport = &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
		}
	}

	client, err := elasticsearch8.NewClient(elasticsearch8.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
		Transport: transport,
	})
	if err != nil {
		return nil, fmt.Errorf("Error creating the client: %s", err)
	}

	return &ElasticsearchClient{client: client}, nil
}

// CreateIndex 创建索引
func (es *ElasticsearchClient) CreateIndex(index string, mapping string) error {
	res, err := es.client.Indices.Create(index, es.client.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		return fmt.Errorf("Error creating index: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}
	return nil
}

// AddDocument 添加文档
func (es *ElasticsearchClient) AddDocument(index string, id string, doc interface{}) error {
	docJson, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("Error marshalling document: %s", err)
	}

	res, err := es.client.Index(index, strings.NewReader(string(docJson)), es.client.Index.WithDocumentID(id))
	if err != nil {
		return fmt.Errorf("Error indexing document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}
	return nil
}

// GetDocument 获取文档
func (es *ElasticsearchClient) GetDocument(index string, id string) (map[string]interface{}, error) {
	res, err := es.client.Get(index, id)
	if err != nil {
		return nil, fmt.Errorf("Error getting document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}

	var doc map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}
	return doc, nil
}

// Search 查询
func (es *ElasticsearchClient) Search(index string, query string) (map[string]interface{}, error) {
	res, err := es.client.Search(func(request *esapi.SearchRequest) {

		request.Index = []string{index}
		request.Body = strings.NewReader(query)
		request.TrackTotalHits = true
	})
	if err != nil {
		return nil, fmt.Errorf("Error search document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}
	var doc map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}
	return doc, nil
}

// SearchWithPagination 查询带分页
func (es *ElasticsearchClient) SearchWithPagination(index string, query map[string]interface{}, from, size int) (map[string]interface{}, error) {
	queryJson, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling query: %s", err)
	}
	res, err := es.client.Search(es.client.Search.WithIndex(index),
		es.client.Search.WithBody(strings.NewReader(string(queryJson))),
		es.client.Search.WithFrom(from),
		es.client.Search.WithSize(size),
	)
	if err != nil {
		return nil, fmt.Errorf("Error executing search document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}
	return result, nil
}

// SearchWithScroll 使用Scroll API分页查询文档(适用于需要遍历大量数据的场景，它通过保持一个快照来避免每次查询重新计算的开销)
func (es *ElasticsearchClient) SearchWithScroll(index string, query map[string]interface{}, scrollTime time.Duration, size int) ([][]byte, error) {
	queryJson, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling query: %s", err)
	}

	var results [][]byte
	res, err := es.client.Search(
		es.client.Search.WithIndex(index),
		es.client.Search.WithBody(strings.NewReader(string(queryJson))),
		es.client.Search.WithScroll(scrollTime),
		es.client.Search.WithSize(size),
	)
	if err != nil {
		return nil, fmt.Errorf("Error executing search query: %s", err)
	}
	defer res.Body.Close()

	for {
		var result map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("Error parsing the response body: %s", err)
		}

		// 收集搜索结果
		hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
		for _, hit := range hits {
			doc, err := json.Marshal(hit)
			if err != nil {
				return nil, fmt.Errorf("Error marshalling hit: %s", err)
			}
			results = append(results, doc)
		}

		// 如果没有更多结果，退出循环
		if len(hits) == 0 {
			break
		}

		// 获取scroll_id
		scrollID := result["_scroll_id"].(string)
		res, err = es.client.Scroll(
			es.client.Scroll.WithScrollID(scrollID),
			es.client.Scroll.WithScroll(scrollTime),
		)
		if err != nil {
			return nil, fmt.Errorf("Error executing scroll query: %s", err)
		}
		defer res.Body.Close()
	}
	return results, nil
}

//// 更新索引设置禁用 _id 字段的 Fielddata
//func (es *ElasticsearchClient) UpdateIndexSettings(index string) error {
//	// 准备要更新的索引设置
//	updateSettings := `{
//		"index": {
//			"id_field_data": {
//				"enabled": false
//			}
//		}
//	}`
//
//	// 调用 Indices.PutSettings 方法
//	req := esapi.IndicesPutSettingsRequest{
//		Index: []string{index}, // 要更新的索引名称
//		Body:  strings.NewReader(updateSettings),
//	}
//
//	res, err := req.Do(context.Background(), es.client.Transport)
//	if err != nil {
//		log.Fatalf("Error updating index settings: %s", err)
//	}
//	defer res.Body.Close()
//
//	if res.IsError() {
//		log.Fatalf("Error updating index settings: %s", res.Status())
//	}
//
//	fmt.Println("Index settings updated successfully")
//	return nil
//}
//
//// SearchWithSearchAfter 使用Search After分页查询文档(适用于深度分页，它基于排序字段来避免深度分页的性能问题)
//
///*
//注意 禁用 _id 字段的 Fielddata
//*/
//func (es *ElasticsearchClient) SearchWithSearchAfter(index string, query map[string]interface{}, sortFields []string, size int) ([][]byte, error) {
//	queryJson, err := json.Marshal(query)
//	if err != nil {
//		return nil, fmt.Errorf("Error marshalling query: %s", err)
//	}
//
//	var results [][]byte
//	var searchAfter []interface{}
//
//	for {
//		var result map[string]interface{}
//		var res *esapi.Response
//
//		// 包含 search_after 字段
//		if len(searchAfter) > 0 {
//			query["search_after"] = searchAfter
//		}
//
//		queryJson, err = json.Marshal(query)
//		if err != nil {
//			return nil, fmt.Errorf("Error marshalling query: %s", err)
//		}
//
//		res, err = es.client.Search(
//			es.client.Search.WithIndex(index),
//			es.client.Search.WithBody(strings.NewReader(string(queryJson))),
//			es.client.Search.WithSort(sortFields...),
//			es.client.Search.WithSize(size),
//		)
//		if err != nil {
//			return nil, fmt.Errorf("Error executing search query: %s", err)
//		}
//		defer res.Body.Close()
//
//		if res.IsError() {
//			return nil, fmt.Errorf("Error response from Elasticsearch: %s", res.String())
//		}
//
//		if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
//			return nil, fmt.Errorf("Error parsing the response body: %s", err)
//		}
//
//		// 收集搜索结果
//		hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
//		for _, hit := range hits {
//			doc, err := json.Marshal(hit)
//			if err != nil {
//				return nil, fmt.Errorf("Error marshalling hit: %s", err)
//			}
//			results = append(results, doc)
//		}
//
//		// 如果没有更多结果，退出循环
//		if len(hits) == 0 {
//			break
//		}
//
//		// 获取最后一个文档的排序值
//		lastHit := hits[len(hits)-1].(map[string]interface{})
//		searchAfter = lastHit["sort"].([]interface{})
//	}
//	return results, nil
//}

// UpdateDocument 更新文档
func (es *ElasticsearchClient) UpdateDocument(index string, id string, doc interface{}) error {
	updateDoc := map[string]interface{}{
		"doc": doc,
	}
	docJson, err := json.Marshal(updateDoc)
	if err != nil {
		return fmt.Errorf("Error marshalling document: %s", err)
	}

	res, err := es.client.Update(index, id, strings.NewReader(string(docJson)))
	if err != nil {
		return fmt.Errorf("Error updating document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}
	return nil
}

//func (es *ElasticsearchClient) UpdateDocument(index string, id string, doc interface{}) error {
//	docJson, err := json.Marshal(doc)
//	if err != nil {
//		return fmt.Errorf("Error marshalling document: %s", err)
//	}
//
//	res, err := es.client.Update(index, id, strings.NewReader(string(docJson)))
//	if err != nil {
//		return fmt.Errorf("Error updating document: %s", err)
//	}
//	defer res.Body.Close()
//	if res.IsError() {
//		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
//	}
//	return nil
//}

// DeleteDocument 删除文档
func (es *ElasticsearchClient) DeleteDocument(index string, id string) error {
	res, err := es.client.Delete(index, id)
	if err != nil {
		return fmt.Errorf("Error deleting document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}
	return nil
}
