package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"github.com/richelieu42/chimera/v2/src/web/ginKit"
	"net/http"
)

func main() {
	config := &ginKit.Config{
		Port:       80,
		Colorful:   true,
		Middleware: nil,
	}
	ginKit.MustSetUp(config, nil, func(engine *gin.Engine) error {
		engine.MaxMultipartMemory = 8 << 20

		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
		engine.Any("/upload", func(ctx *gin.Context) {
			// 单文件
			file, err := ctx.FormFile("file")
			if err != nil {
				ctx.String(http.StatusOK, err.Error())
				return
			}

			// 上传文件到指定的路径
			// ctx.SaveUploadedFile(file, dst)
			ctx.String(http.StatusOK, fmt.Sprintf("'%s'(%s) uploaded!", file.Filename, dataSizeKit.ToReadableStringWithIEC(uint64(file.Size))))
		})
		return nil
	})
}
