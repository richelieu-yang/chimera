package ginKit

import (
	"github.com/gin-gonic/gin"
)

// NewEngine
/*
@param recovery 可以为nil，将采用默认值 gin.Recovery()
*/
func NewEngine() *gin.Engine {
	engine := gin.New()

	// 默认: []string{"X-Forwarded-For", "X-Real-IP"}
	engine.RemoteIPHeaders = []string{"X-Real-IP", "Client-IP", "X-Forwarded-For"}

	// 默认: true
	engine.RedirectTrailingSlash = true

	return engine
}
