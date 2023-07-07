package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	type config struct {
		Gin *Config `json:"Gin"`
	}

	wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}
	logrus.Infof("wd: [%s].", wd)

	c := &config{}
	path := "chimera-lib/config.yaml"
	confKit.MustLoad(path, c)
	MustSetUp(c.Gin, nil, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			//if err := httpKit.Proxy(ctx.Writer, ctx.Request, "http", "127.0.0.1:8080"); err != nil {
			//	ctx.ToDSN(http.StatusInternalServerError, err.Error())
			//	return
			//}

			ctx.String(http.StatusOK, "hello")
		})
		return nil
	})
}
