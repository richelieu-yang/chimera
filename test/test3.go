package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
)

func main() {
	engine := gin.Default()

	ginKit.DefaultFavicon(engine)
	if err := ginKit.DefaultNoRoute(engine); err != nil {
		panic(err)
	}
	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, "ok")
		ctx.Status(500)
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
