package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/web/ginKit"
	"net/http"
)

func main() {
	config := &ginKit.Config{
		Port:       8888,
		Colorful:   true,
		Middleware: nil,
	}
	ginKit.MustSetUp(config, nil, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			str := ginKit.ObtainParam(ctx, "cyy")

			ctx.String(http.StatusOK, str)
		})
		return nil
	})
}
