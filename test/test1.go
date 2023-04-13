package main

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/web/ginKit"
	"io"
	"net/http"
	"os"
)

func main() {
	os.File{}
	io.SeekStart
	io.SeekCurrent
	io.SeekEnd

	io.Seeker()

	bytes.NewBuffer()
	bytes.NewBufferString()

	http.MaxBytesReader()

	bufio.Reader{}

	bytes.Buffer{}

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
