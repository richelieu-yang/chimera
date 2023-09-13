package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		fmt.Println(ctx.Request.Body == http.NoBody)

		fmt.Println(ctx.PostForm("a"))
		fmt.Println(ctx.PostFormArray("a"))
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
