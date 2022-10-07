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

	//限制上传最大尺寸
	engine.MaxMultipartMemory = 16 << 20

	engine.RedirectTrailingSlash = true

	return engine
}
