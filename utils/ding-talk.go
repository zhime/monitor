package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/zhime/monitor/global"
)

// DingTalkMessage 消息结构体
type DingTalkMessage struct {
	MsgType string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

// 生成签名
func GenerateSignature(secret string, timestamp int64) (string, error) {
	message := fmt.Sprintf("%d\n%s", timestamp, secret)
	mac := hmac.New(sha256.New, []byte(secret))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	signature := mac.Sum(nil)
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	return url.QueryEscape(encodedSignature), nil
}

// SendDingTalkNotification 发送钉钉通知
func SendDingTalkNotification(message string) error {
	timestamp := time.Now().Unix() * 1000

	signature, err := GenerateSignature(global.CONFIG.DingTalk.Secret, timestamp)
	if err != nil {
		return fmt.Errorf("生成签名失败: %v", err)
	}

	// 创建消息
	msg := DingTalkMessage{
		MsgType: "text",
	}
	msg.Text.Content = message

	// 将消息转换为 JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("JSON marshal error: %v", err)
	}

	// 构造完整的请求 URL
	requestURL := fmt.Sprintf("%s&timestamp=%d&sign=%s", global.CONFIG.DingTalk.Webhook, timestamp, signature)

	// 发送 HTTP POST 请求
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("HTTP POST error: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
