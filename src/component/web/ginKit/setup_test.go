package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	{
		logrusKit.MustSetUp(nil)

		if err := pathKit.SetTempDir("_temp"); err != nil {
			panic(err)
		}

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
		//engine.Any("*path", func(ctx *gin.Context) {
		//	if err := proxyKit.Proxy(ctx.Writer, ctx.Request, "127.0.0.1:8888"); err != nil {
		//		ctx.String(500, err.Error())
		//		return
		//	}
		//})

		err := RegisterHandlers(engine, "test", []string{http.MethodHead, http.MethodPost}, func(ctx *gin.Context) {
			keys := []string{
				"Host",
				"X-Real-IP",
				"Client-IP",
				"X-Forwarded-For",
				"X-Forwarded-Proto",
			}
			for _, key := range keys {
				s := httpKit.GetHeaderValues(ctx.Request.Header, key)
				logrus.Infof("%s: %s", key, s)
			}
			logrus.Info("======")

			ctx.String(200, "test")
		})
		if err != nil {
			return err
		}

		//engine.Any("/a/b", func(ctx *gin.Context) {
		//	ctx.String(200, "hello world")
		//})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(true))
}
