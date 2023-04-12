package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"github.com/richelieu42/chimera/v2/src/web/ginKit"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	bytes.NewReader()
	strings.NewReader()
	bytes.NewReader()
	os.File

	bytes.Buffer{}

	io.Seeker()
	http.MaxBytesReader()

	config := &ginKit.Config{
		Port:       80,
		Colorful:   true,
		Middleware: nil,
	}
	ginKit.MustSetUp(config, nil, func(engine *gin.Engine) error {
		engine.Use(func(ctx *gin.Context) {
			ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, int64(30<<20))
		})

		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
		engine.Any("/upload", func(ctx *gin.Context) {
			//// 多文件上传
			//form, err := ctx.MultipartForm()
			//if err != nil {
			//	ctx.String(http.StatusOK, err.Error())
			//	return
			//}
			//form = form

			// 单文件上传
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
