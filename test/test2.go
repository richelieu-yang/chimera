package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		a := ctx.PostForm("a")
		fmt.Printf("a: [%s].\n", a)
		ctx.String(http.StatusOK, a)
	})

	if err := engine.Run(":12345"); err != nil {
		panic(err)
	}
}
