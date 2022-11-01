package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		referer := ctx.GetHeader("referer")
		ctx.String(http.StatusOK, fmt.Sprintf("referer: [%s]", referer))
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
