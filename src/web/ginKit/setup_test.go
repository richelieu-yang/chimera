package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/richelieu42/chimera/v2/src/core/osKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"net/http"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	logrusKit.MustSetUp(nil)

	var path string
	if osKit.IsWindows() {
		path = "D:\\GolandProjects\\chimera\\chimera-lib\\config.yaml"
	} else {
		path = "/Users/richelieu/GolandProjects/chimera/chimera-lib/config.yaml"
	}

	type config struct {
		Gin *Config `json:"Gin"`
	}
	c := &config{}
	confKit.MustLoad(path, c)
	MustSetUp(c.Gin, nil, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			//if err := httpKit.Proxy(ctx.Writer, ctx.Request, "http", "127.0.0.1:8080"); err != nil {
			//	ctx.String(http.StatusInternalServerError, err.Error())
			//	return
			//}
			ctx.String(http.StatusOK, "hello")
		})
		return nil
	})
}

//func TestMustSetUp1(t *testing.T) {
//	config := &Config{
//		Port:       80,
//		Colorful:   true,
//		Middleware: nil,
//		SSL: &SslConfig{
//			CertFile: "/Users/richelieu/GolandProjects/chimera/chimera-lib/ssl.pem",
//			KeyFile:  "/Users/richelieu/GolandProjects/chimera/chimera-lib/ssl.key",
//			Port:     443,
//		},
//	}
//	MustSetUp(config, nil, func(engine *gin.Engine) error {
//		engine.Any("/test", func(ctx *gin.Context) {
//			ctx.String(http.StatusOK, "ok")
//		})
//		return nil
//	})
//}
