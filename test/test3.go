package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		str, err := ioKit.ReadStringFromReader(ctx.Request.Body)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(str)
		ctx.String(http.StatusOK, str)
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
