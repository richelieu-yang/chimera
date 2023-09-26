package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	port := 8500

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		logrus.Fatal(err)
	}
}
