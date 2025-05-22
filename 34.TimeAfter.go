package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 创建一个请求
	req, err := http.NewRequest("GET", "https://www.example.com", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 创建一个通道，用于接收请求结果
	resultCh := make(chan *http.Response, 1)

	// 启动一个 Goroutine 发送请求
	go func() {
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			resultCh <- nil // 发送 nil 表示请求失败
			return
		}
		resultCh <- resp // 发送响应结果
	}()

	// 使用 select 监听请求结果和超时,

	// ！！！！！！！！！优雅，太优雅了 goroutine 异步请求api,select + case处理rsp和超时
	select {
	case resp := <-resultCh:
		if resp != nil {
			fmt.Println("Response received:", resp.Status)
			defer resp.Body.Close()
		} else {
			fmt.Println("Request failed")
		}
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout: Request took too long")
	}
}
