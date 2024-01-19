package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/richelieu-yang/chimera/v2/src/log/logrusInitKit"
)

func main() {
	engine := gin.Default()

	engine.Any("/", func(ctx *gin.Context) {
		ctx.String(200, "hello world")
	})

	if err := engine.Run(":8888"); err != nil {
		panic(err)
	}
}
