package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()

	engine.Any("/ping", func(ctx *gin.Context) {
		time.Sleep(time.Second * 3)
		ctx.String(http.StatusOK, "ojbk")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
