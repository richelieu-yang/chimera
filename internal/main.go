package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/httpKit"
	_ "github.com/richelieu-yang/chimera/v2/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
)

func main() {
	engine := gin.Default()

	engine.Any("/", func(ctx *gin.Context) {
		logrus.Infof("client shceme: [%s]", httpKit.GetClientScheme(ctx.Request))
		ctx.String(200, "hello world")
	})
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(404, "no route")
	})

	if err := engine.Run(":8888"); err != nil {
		panic(err)
	}
}
