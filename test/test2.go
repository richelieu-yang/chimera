package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"github.com/richelieu42/chimera/v2/src/web/ginKit"
	"github.com/richelieu42/chimera/v2/src/web/httpKit"
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
		engine.POST("/test", func(ctx *gin.Context) {
			err := httpKit.MakeRequestBodySeekable(ctx.Request)
			if err != nil {
				ctx.String(http.StatusInternalServerError, err.Error())
				return
			}

			str := ginKit.ObtainParam(ctx, "cyy")
			logrus.Info(str)

			err = httpKit.Proxy(ctx.Writer, ctx.Request, nil, "http", "127.0.0.1:8888", nil, nil)
			if err != nil {
				ctx.String(http.StatusInternalServerError, err.Error())
				return
			}
		})

		engine.POST("/upload", func(ctx *gin.Context) {
			file, err := ctx.FormFile("file")
			if err != nil {
				if mbe, ok := err.(*http.MaxBytesError); ok {
					// http状态码: 413（由于请求实体过大，服务器无法处理，因此拒绝请求）
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
