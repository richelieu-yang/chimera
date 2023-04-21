package main

import (
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"log"
	"net/http"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

func handler(ctx *gin.Context) {
	_ = ctx.PostForm("b")
	if len(ctx.Errors) > 0 {
		return
	}
	ctx.String(http.StatusOK, dataSizeKit.ToReadableStringWithIEC(uint64(ctx.Request.ContentLength)))
}

func main() {
	r := gin.Default()
	r.Use(limits.RequestSizeLimiter(10))
	r.POST("/", handler)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
