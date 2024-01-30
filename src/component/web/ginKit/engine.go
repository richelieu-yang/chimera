package ginKit

import (
	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	engine := gin.New()

	// 默认: []string{"X-Forwarded-For", "X-Real-IP"}
	engine.RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP", "Client-IP"}

	// 默认: true
	engine.RedirectTrailingSlash = true

	return engine
}
