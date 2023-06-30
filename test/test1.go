package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Any("/*.act", func(ctx *gin.Context) {
		logrus.Infof("ctx.FullPath(): [%s]", ctx.FullPath())
		logrus.Infof("ctx.Param(\"a\"): [%s]", ctx.Param(".act"))

		ctx.String(http.StatusOK, ctx.FullPath())
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
