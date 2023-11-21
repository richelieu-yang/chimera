package ginKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}
	logrus.Infof("wd: [%s].", wd)

	path := "_chimera-lib/config.yaml"

	type config struct {
		Gin *Config `json:"Gin"`
	}
	c := &config{}
	_, err = viperKit.UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}

	MustSetUp(c.Gin, nil, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			//if err := httpKit.Proxy(ctx.Writer, ctx.Request, "http", "127.0.0.1:8080"); err != nil {
			//	ctx.ToDsnString(http.StatusInternalServerError, err.Error())
			//	return
			//}

			//ctx.Bind()
			//ctx.ShouldBind()
			//
			//time.Sleep(time.Second * 3)

			//ctx.String(http.StatusOK, "hello")

			fmt.Println(ctx.Request.Header["Content-Type"])

			//fmt.Println(ioKit.ReadStringFromReader(ctx.Request.Body))

			fmt.Println(ObtainParam(ctx, "name"))

			ctx.String(200, "hello")
		})
		return nil
	})
}
