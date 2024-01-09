package main

import (
	"fmt"
	"github.com/CAFxX/httpcompression"
	"net/http"
)

func main() {
	// 创建一个 HTTP 处理器
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!")
	})

	// 获取一个压缩适配器
	compress, err := httpcompression.DefaultAdapter()
	if err != nil {
		panic(err)
	}

	// 注册一个路由，使用压缩适配器包装处理器
	http.Handle("/hello", compress(handler))

	// 启动 HTTP 服务器
	http.ListenAndServe(":8080", nil)
}
