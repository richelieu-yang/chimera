package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
)

func main() {
	logrusInitKit.SetUp()
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {

	})

	if err := engine.Run(":10000"); err != nil {
		panic(err)
	}
}
