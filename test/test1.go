package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()

	engine.Use(func(ctx *gin.Context) {
		if true {
			e := ctx.AbortWithError(http.StatusNotExtended, redis.Nil)
			fmt.Println(e.Error())
			return
		}
	})

	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, time.Now().String())
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
