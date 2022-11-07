package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/ginKit"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.New()

	middleware, err := ginKit.NewRefererMiddleware(true, false, "", "*.yozo.com")
	if err != nil {
		panic(err)
	}
	r.Use(gin.Logger(), gin.Recovery(), middleware)

	r.Use(func(ctx *gin.Context) {
		/* 第1个自定义中间件 */
		logrus.Info("[A] before")
		ctx.Next()
		logrus.Info("[A] after")
	}, func(ctx *gin.Context) {
		/* 第2个自定义中间件 */
		logrus.Info("[B] before")
		ctx.Next()
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
