package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	// 设置NoRoute处理器
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "页面不存在"})
	})

	// 设置NoMethod处理器
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(func(ctx *gin.Context) {
		var allowed []string
		route := ctx.Request.URL.Path
		routeInfoSlice := engine.Routes()
		for _, routeInfo := range routeInfoSlice {
			if routeInfo.Path == route {
				allowed = append(allowed, routeInfo.Method)
			}
		}

		text := fmt.Sprintf("Method(%s) isn't allowed for route(%s) and allowed methods is %s.", ctx.Request.Method, ctx.Request.URL.Path, allowed)
		ctx.String(http.StatusMethodNotAllowed, text)
	})

	// 其他路由和方法
	engine.POST("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
