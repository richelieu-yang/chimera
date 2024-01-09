package ginKit

import (
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

	MustSetUp(c.Gin, nil, func(engine *gin.Engine) error {
		engine.Any("a/*path", func(ctx *gin.Context) {
			ctx.String(200, ctx.Param("path"))
		})
		//engine.Any("a/:path", func(ctx *gin.Context) {
		//	ctx.String(200, ctx.Param("path"))
		//})

		engine.NoRoute(func(ctx *gin.Context) {
			ctx.String(404, "[TARGET] no route")
		})
		return nil
	}, "TEST")
}
