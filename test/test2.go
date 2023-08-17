package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()

	engine.Any("/ping", func(ctx *gin.Context) {
		time.Sleep(time.Second * 60)
		ctx.String(http.StatusOK, "OK")
	})

	if err := engine.Run(":80"); err != nil {
		panic(engine)
	}
}
