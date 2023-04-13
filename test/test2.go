package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/richelieu42/chimera/v2/src/web/ginKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

type config struct {
	Gin *ginKit.Config
}

func main() {
	c := &config{}
	confKit.MustLoad("/Users/richelieu/GolandProjects/chimera/chimera-lib/env.yaml", c)

	ginKit.MustSetUp(c.Gin, nil, func(engine *gin.Engine) error {
		engine.Use(func(ctx *gin.Context) {
			logrus.Info("before Next()")
			ctx.Next()
			logrus.Info("after Next()")
		}, func(ctx *gin.Context) {
			logrus.Info("before Abort()")
			ctx.AbortWithError(http.StatusInternalServerError, redis.Nil)
			logrus.Info("after Abort()")
		})

		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
		return nil
	})
}
