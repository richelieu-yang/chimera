package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
)

func NewEngine() *gin.Engine {
	engine := gin.New()

	// 默认: true
	engine.ForwardedByClientIP = true
	// 默认: []string{"X-Forwarded-For", "X-Real-IP"}
	engine.RemoteIPHeaders = httpKit.RemoteIPHeaders

	// 默认: true
	engine.RedirectTrailingSlash = true

	return engine
}
