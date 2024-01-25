package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
)

func main() {
	engine := gin.Default()

	engine.Use(ginKit.NewOptionsMiddleware())

	engine.GET("/test", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

	if err := engine.Run(":80"); err != nil {
		fmt.Println(err)
	}
}
