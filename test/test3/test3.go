package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 添加一个全局中间件来设置HSTS响应头
	r.Use(func(c *gin.Context) {
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the secure server!",
		})
	})

	// 启动服务（这里假设是HTTPS服务）
	err := r.RunTLS(":443", "_chimera-lib/ssl.crt", "_chimera-lib/ssl.key")
	if err != nil {
		panic(err)
	}
}
