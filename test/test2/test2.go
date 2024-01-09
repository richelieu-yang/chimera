package main

import (
	"fmt"
	"github.com/CAFxX/httpcompression"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	// 获取一个压缩适配器
	compress, err := httpcompression.DefaultAdapter()
	if err != nil {
		panic(err)
	}
	compress(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(666)
	}))

	engine.Use(func(ctx *gin.Context) {

		compress(ctx.Writer, ctx.Request)
		ctx.Next()
	})

	engine.Any("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!Hello, world!")
	})

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
