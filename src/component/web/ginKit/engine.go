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

	engine.RedirectTrailingSlash = true

	return engine
}
