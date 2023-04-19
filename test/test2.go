package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	engine := gin.Default()

	if err := engine.SetTrustedProxies([]string{"192.168.9.254"}); err != nil {
		logrus.Fatal(err)
	}

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is 8001")
	})

	if err := engine.Run(":8001"); err != nil {
		logrus.Fatal(err)
	}
}
