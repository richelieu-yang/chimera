package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	engine := gin.New()
	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ":8001")
	})
	if err := engine.Run(":8001"); err != nil {
		logrus.Fatal(err)
	}
}
