package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	_ "github.com/richelieu-yang/chimera/v2/src/log/logrusInitKit"
)

func main() {
	engine := gin.Default()

	engine.Use(ginKit.NewCorsMiddleware(nil))

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, "hello world")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
