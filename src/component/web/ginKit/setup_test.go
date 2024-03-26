package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/proxyKit"
	"github.com/richelieu-yang/chimera/v3/src/config/yaml/yamlKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	{
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
		Gin *Config `json:"gin" yaml:"gin"`
	}
	c := &config{}
	/*
		TODO: 反序列化.yaml文件，先用 yamlKit.UnmarshalFromFile 替换 viperKit.UnmarshalFromFile，原因: https://github.com/spf13/viper/issues/1769
	*/
	err := yamlKit.UnmarshalFromFile(path, c)
	//_, err := viperKit.UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}

	MustSetUp(c.Gin, func(engine *gin.Engine) error {
		engine.Use(func(ctx *gin.Context) {
			ctx.String(333, "three three three")
			ctx.Abort()
			return
		})

		engine.Any("/test", func(ctx *gin.Context) {
			qm := map[string][]string{
				"b": {"bOx"},
				"c": {"阿德去外地"},
			}

			if err := proxyKit.ProxyWithGin(ctx, "127.0.0.1:10000", proxyKit.WithExtraQueryParams(qm)); err != nil {
				ctx.String(http.StatusInternalServerError, err.Error())
				return
			}
		})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(true))
}
