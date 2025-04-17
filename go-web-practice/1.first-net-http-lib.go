package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func helloHander(rsp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rsp, "Hello, %q\n", html.EscapeString(req.URL.Path))
}

func headerHandler(rsp http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(rsp, "%v:%v\n", name, h)
		}
	}
}

// 定义处理器函数，注册处理器函数，启动服务器
func main() {
	http.HandleFunc("/hello", helloHander)
	http.HandleFunc("/headers", headerHandler)

	fmt.Println("Start sever at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
