package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/resources"
	"net/http"
)

func main() {
	engine := gin.Default()

	fs := resources.AssetFile()
	engine.StaticFS("/s", fs)

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, fileKit.GetFileName("404.html"), gin.H{
			"urlPath": ctx.Request.URL.Path,
		})
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
