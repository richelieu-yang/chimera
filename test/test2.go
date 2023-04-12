package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"github.com/richelieu42/chimera/v2/src/web/ginKit"
	"net/http"
)

type config struct {
	Gin *ginKit.Config
}

func main() {
	c := &config{}
	confKit.MustLoad("/Users/richelieu/GolandProjects/chimera/chimera-lib/env.yaml", c)

	ginKit.MustSetUp(c.Gin, nil, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})

		engine.POST("/upload", func(ctx *gin.Context) {
			file, err := ctx.FormFile("file")
			if err != nil {
				if mbe, ok := err.(*http.MaxBytesError); ok {
					ctx.String(http.StatusRequestEntityTooLarge, mbe.Error())
					return
				}
				ctx.String(http.StatusOK, err.Error())
				return
			}

			ctx.String(http.StatusOK, fmt.Sprintf("'%s'(%s) uploaded!", file.Filename, dataSizeKit.ToReadableStringWithIEC(uint64(file.Size))))
		})

		return nil
	})
}
