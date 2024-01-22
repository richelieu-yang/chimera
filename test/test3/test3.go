package main

import (
	"github.com/gin-gonic/gin"
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
		ctx.JSON(405, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "方法不允许"})
	})

	// 其他路由和方法
	engine.POST("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
