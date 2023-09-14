package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	if err := engine.Run(":8888"); err != nil {
		panic(err)
	}
}
