package ginKit

import "github.com/gin-gonic/gin"

type Handler func(ctx *gin.Context) (*ResponsePackage, error)

// WrapToHandlerFunc Handler 类型 => gin.HandlerFunc 类型
func WrapToHandlerFunc(handler Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if handler == nil {
			return
		}

		pack, err := handler(ctx)
		RespondPackageOrError(ctx, pack, err)
	}
}
