package ginKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	config := &Config{
		Port:       8000,
		Colorful:   true,
		Middleware: nil,
		SSL: &SslConfig{
			CertFile: "/Users/richelieu/GolandProjects/chimera/ssl.pem",
			KeyFile:  "/Users/richelieu/GolandProjects/chimera/ssl.key",
			Port:     443,
		},
	}
	MustSetUp(config, nil, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
		return nil
	})
}
