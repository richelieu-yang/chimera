package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	_ "github.com/richelieu-yang/chimera/v2/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
)

func main() {
	engine := gin.Default()

	engine.Use(ginKit.NewCorsMiddleware(nil))

	engine.POST("/test", func(ctx *gin.Context) {
		data, _ := ioKit.ReadFromReader(ctx.Request.Body)
		logrus.Infof("data: [%s]", data)

		ctx.String(200, "hello world")
	})
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(404, "no route")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
