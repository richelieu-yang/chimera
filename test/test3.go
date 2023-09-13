package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

func main() {
	engine := gin.Default()

	engine.POST("/test", func(ctx *gin.Context) {
		_, ok := ctx.Request.Body.(io.Seeker)
		fmt.Println(ok)

		//fmt.Println(ctx.PostForm("a"))
		//fmt.Println(ctx.PostFormArray("a"))
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
