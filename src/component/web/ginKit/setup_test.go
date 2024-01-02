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

	MustSetUp(c.Gin, "service info", nil, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(200, "ok")

			//defer func() {
			//	if obj := recover(); obj != nil {
			//		if err, ok := obj.(error); ok {
			//			if strKit.ContainsIgnoreCase(err.Error(), "net/http: abort Handler") {
			//				// 忽略
			//				return
			//			}
			//		}
			//		logrus.WithField("obj", obj).Error("panic")
			//		debug.PrintStack()
			//	}
			//}()
			//
			//if err := httpKit.Proxy(ctx.Writer, ctx.Request, "127.0.0.1:8888", httpKit.WithReqUrlPath(ptrKit.Of("/push/sse"))); err != nil {
			//	logrus.Error(err)
			//	return
			//}
			//logrus.Info("Manager to proxy.")
		})
		return nil
	})
}
