package es

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestHttpAdd(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}

	// 示例：添加文档
	doc := map[string]interface{}{
		"title":   "Test DocumentBBB",
		"content": "This is a test documentBB.",
	}
	err = esClientHTTP.AddDocument("test-index", "1", doc)
	if err != nil {
		log.Fatalf("Error adding document (HTTP): %s", err)
	}
	log.Println("Document added successfully (HTTP)")
}
func TestHttpGet(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	// 示例：获取文档
	docRetrieved, err := esClientHTTP.GetDocument("test-index", "1")
	if err != nil {
		log.Fatalf("Error getting document (HTTP): %s", err)
	}
	log.Printf("Document retrieved (HTTP): %+v\n", docRetrieved)
}
func TestHttpUpdate(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	// 示例：更新文档
	docUpdate := map[string]interface{}{
		"title":   "Updated Document",
		"content": "This is an updated test document.",
	}
	err = esClientHTTP.UpdateDocument("test-index", "1", docUpdate)
	if err != nil {
		log.Fatalf("Error updating document (HTTP): %s", err)
	}
	log.Println("Document updated successfully (HTTP)")
}
func TestHttpDelete(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	// 示例：删除文档
	err = esClientHTTP.DeleteDocument("test-index", "1")
	if err != nil {
		log.Fatalf("Error deleting document (HTTP): %s", err)
	}
	log.Println("Document deleted successfully (HTTP)")
}

// 查询带分页
func TestSearchPaging(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	// 查询条件
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	// 分页查询
	from := 0
	size := 10

	result, err := esClientHTTP.SearchWithPagination("test-index", query, from, size)
	if err != nil {
		log.Fatalf("Error executing search query: %s", err)
	}

	fmt.Printf("Search results: %+v\n", result)
}

// 大数据量分页A
func TestSearchWithScroll(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	// 查询条件
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	// 分页查询
	scrollTime := 2 * time.Minute
	size := 100

	results, err := esClientHTTP.SearchWithScroll("test-index", query, scrollTime, size)
	if err != nil {
		log.Fatalf("Error executing search query: %s", err)
	}

	for _, result := range results {
		fmt.Printf("Document: %s\n", result)
	}
}

//	func TestDisableId(t *testing.T) {
//		// 使用HTTP
//		esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
//		if err != nil {
//			log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
//		}
//
//		err = esClientHTTP.UpdateIndexSettings("test-index")
//		fmt.Println(err)
//	}
//
// // 大数据量分页B
//
//	func TestSearchWithSearchAfter(t *testing.T) {
//		// 使用HTTP
//		esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
//		if err != nil {
//			log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
//		}
//		// 查询条件
//		query := map[string]interface{}{
//			"query": map[string]interface{}{
//				"match_all": map[string]interface{}{},
//			},
//		}
//
//		// 分页查询
//		sortFields := []string{"_id"} // 使用文档ID排序
//		size := 100
//
//		results, err := esClientHTTP.SearchWithSearchAfter("test-index", query, sortFields, size)
//		if err != nil {
//			log.Fatalf("Error executing search query: %s", err)
//		}
//
//		for _, result := range results {
//			fmt.Printf("Document: %s\n", result)
//		}
//	}
func TestES(test *testing.T) {
	// 使用HTTPS和证书
	esClientTLS, err := NewElasticsearchClient(true, "D:\\other\\lxgame\\ese\\CloudSearchService.cer", "admin", "123456", []string{"https://21.1.22.10:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (TLS): %s", err)
	}

	// 示例：创建索引
	indexMapping := `{
		"mappings": {
			"properties": {
				"title": { "type": "text" },
				"content": { "type": "text" }
			}
		}
	}`
	err = esClientTLS.CreateIndex("test-index", indexMapping)
	if err != nil {
		log.Fatalf("Error creating index (TLS): %s", err)
	}
	log.Println("Index created successfully (TLS)")

}
