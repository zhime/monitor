package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

func main() {
	// 配置 Elasticsearch 客户端
	cfg := elasticsearch8.Config{
		Addresses: []string{
			"https://192.168.0.200:9200",
		},
		Username: "elastic",
		Password: "elastic",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// 创建 Elasticsearch 客户端
	es, err := elasticsearch8.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 测试连接 - 打印集群信息
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	fmt.Println(res)
}
