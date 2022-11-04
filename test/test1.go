package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(func(ctx *gin.Context) {
		/* 第1个自定义中间件 */
		logrus.Info("[A] before")
		ctx.Next()
		logrus.Info("[A] after")
	}, func(ctx *gin.Context) {
		/* 第2个自定义中间件 */
		logrus.Info("[B] before")
		ctx.Abort()
		logrus.Info("[B] after")
	}, func(ctx *gin.Context) {
		/* 第3个自定义中间件 */
		logrus.Info("[C] before")
		ctx.Next()
		logrus.Info("[C] after")
	})

	r.GET("/test", func(c *gin.Context) {
		logrus.Info("[Handler]")
		c.String(200, "OK")
	})
	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
