package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Any("/*name", func(ctx *gin.Context) {
		str := ctx.Param("name")
		if str != "" {
			fmt.Println(str)
		} else {
			fmt.Println("empty")
		}

		ctx.String(http.StatusOK, "ok")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
