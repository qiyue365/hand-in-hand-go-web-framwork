package main

import (
	"net/http"

	"github.com/qiyue365/hand-in-hand-go-web-framwork/framwork"
)

func main() {
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: framwork.NewCore(),
		// 请求监听地址
		Addr: ":12345",
	}
	server.ListenAndServe()
}
