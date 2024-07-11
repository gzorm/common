package es

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
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
