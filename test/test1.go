package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/web/httpKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	engine := gin.Default()

	//if err := engine.SetTrustedProxies(nil); err != nil {
	//	logrus.Fatal(err)
	//}

	engine.Any("/test", func(ctx *gin.Context) {
		err := httpKit.Proxy(ctx.Writer, ctx.Request, nil, "http", "192.168.60.206:8001", nil, nil)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
	})

	if err := engine.Run(":8000"); err != nil {
		logrus.Fatal(err)
	}
}
