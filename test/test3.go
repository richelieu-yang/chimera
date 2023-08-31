package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var path = "a.txt"

func main() {
	engine := gin.Default()

	engine.Use(func(ctx *gin.Context) {
		logrus.WithFields(logrus.Fields{
			"path":     ctx.Request.URL.Path,
			"rawQuery": ctx.Request.URL.RawQuery,
			"RemoteIP": ctx.RemoteIP(),
			"ClientIP": ctx.ClientIP(),
		}).Info("[DEBUG] Receive a request.")

		ctx.Next()
	})

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ojbk")
	})

	if err := engine.Run(":8888"); err != nil {
		panic(err)
	}
}
