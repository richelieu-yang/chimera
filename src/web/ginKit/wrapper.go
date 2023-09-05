package ginKit

import "github.com/gin-gonic/gin"

type Handler func(ctx *gin.Context) (*ResponsePackage, error)

// WrapToHandlerFunc Handler 类型 => gin.HandlerFunc 类型
/*
@param handler 不能为nil
*/
func WrapToHandlerFunc(handler Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pack, err := handler(ctx)
		RespondPackageOrError(ctx, pack, err)
	}
}
