package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// PrometheusResponse 结构体，用于解析 Prometheus API 响应
type PrometheusResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []interface{}     `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

// 查询 Prometheus API
func queryPrometheus(query string) (*PrometheusResponse, error) {
	url := fmt.Sprintf("http://your-prometheus-server/api/v1/query?query=%s", query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to query Prometheus: %v", err)
	}
	defer resp.Body.Close()

	var result PrometheusResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode Prometheus response: %v", err)
	}
	return &result, nil
}

// 连接 MySQL 并插入数据
func insertIntoMySQL(db *sql.DB, metricName string, value string) error {
	insertQuery := "INSERT INTO prometheus_data (metric_name, value) VALUES (?, ?)"
	_, err := db.Exec(insertQuery, metricName, value)
	if err != nil {
		return fmt.Errorf("failed to insert data into MySQL: %v", err)
	}
	return nil
}

func main() {
	// 连接 MySQL 数据库
	dsn := "your_mysql_user:your_mysql_password@tcp(your_mysql_host:3306)/your_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// 从 Prometheus 查询数据
	query := "up" // 替换为你的 Prometheus 查询
	promData, err := queryPrometheus(query)
	if err != nil {
		log.Fatalf("Failed to query Prometheus: %v", err)
	}

	// 解析 Prometheus 数据并插入 MySQL
	for _, result := range promData.Data.Result {
		metricName := result.Metric["__name__"]
		value := result.Value[1].(string) // 值是字符串类型

		err := insertIntoMySQL(db, metricName, value)
		if err != nil {
			log.Printf("Failed to insert metric %s into MySQL: %v", metricName, err)
		} else {
			fmt.Printf("Inserted metric %s with value %s into MySQL\n", metricName, value)
		}
	}
}
