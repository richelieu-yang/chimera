package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/web/httpKit"
	"net/http"
)

func main() {
	engine := gin.Default()

	key := "ccc"
	engine.Any("/", func(ctx *gin.Context) {
		value := httpKit.GetHeader(ctx.Request.Header, key)
		fmt.Printf("%s: [%s]\n", key, value)

		ctx.String(http.StatusOK, "/")
	})
	engine.Any("/a", func(ctx *gin.Context) {
		value := httpKit.GetHeader(ctx.Request.Header, key)
		fmt.Printf("%s: [%s]\n", key, value)

		ctx.String(http.StatusOK, "/a")
	})
	engine.Any("/b/c", func(ctx *gin.Context) {
		value := httpKit.GetHeader(ctx.Request.Header, key)
		fmt.Printf("%s: [%s]\n", key, value)

		ctx.String(http.StatusOK, "/b/c")
	})

	if err := engine.Run(":81"); err != nil {
		panic(err)
	}
}
