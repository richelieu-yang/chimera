package ginKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	{
		logrusKit.MustSetUp(nil)

		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		logrus.Infof("wd: [%s].", wd)
	}

	path := "_chimera-lib/config.yaml"
	type config struct {
		Gin *Config `json:"Gin"`
	}
	c := &config{}
	_, err := viperKit.UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}

	MustSetUp(c.Gin, func(engine *gin.Engine) error {
		//handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "text/plain")
		//	io.WriteString(w, "Hello, World")
		//})

		engine.Any("*path", func(ctx *gin.Context) {
			str := ctx.Request.Header.Get("Accept-Encoding")
			fmt.Println(str)

			ctx.String(200, ctx.Param("path"))
		})

		//engine.Any("a/:path", func(ctx *gin.Context) {
		//	ctx.String(200, ctx.Param("path"))
		//})
		//engine.NoRoute(func(ctx *gin.Context) {
		//	ctx.String(404, "[TARGET] no route")
		//})
		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(false))
}
