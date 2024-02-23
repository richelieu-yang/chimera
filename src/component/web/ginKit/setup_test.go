package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/config/yaml/yamlKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
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
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(200, "ok")
		})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(false))
}
