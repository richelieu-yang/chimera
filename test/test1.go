package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/web/httpKit"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Any("/*name", func(ctx *gin.Context) {
		httpKit.AddHeader(ctx.Request.Header, "ccc", "q1wEydfvq5wujhdg1biqkhgFYTfy")

		if err := httpKit.Proxy(ctx.Writer, ctx.Request, "http", "127.0.0.1:81"); err != nil {
			ctx.String(http.StatusOK, err.Error())
		}
	})
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
