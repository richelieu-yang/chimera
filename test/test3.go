package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		logrus.Infof("a: [%s]", ginKit.ObtainGetParam(ctx, "a"))
		logrus.Infof("b: [%s]", ginKit.ObtainGetParam(ctx, "b"))
		logrus.Infof("c: [%s]", ginKit.ObtainGetParam(ctx, "c"))

		ctx.String(http.StatusOK, "10000")
	})

	if err := engine.Run(":10000"); err != nil {
		panic(err)
	}
}
