package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/richelieu-yang/chimera/v2/src/log/logrusInitKit"
)

func main() {
	engine := gin.Default()

	engine.Use(gzip.Gzip(gzip.BestSpeed))

	engine.Any("/", func(ctx *gin.Context) {
		ctx.String(200, "hello world")
	})
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(404, "no route")
	})

	if err := engine.Run(":8888"); err != nil {
		panic(err)
	}
}
