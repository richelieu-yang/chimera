package otelKit

import (
	"github.com/gin-gonic/gin"
)

// NewOuterGinMiddleware TODO: 适用于: 外层的服务
/*
PS: 需要先 SetUp || MustSetUp !!!
*/
func NewOuterGinMiddleware() (gin.HandlerFunc, error) {
	return func(ctx *gin.Context) {
		//tracer := otel.Tracer("", trace.WithInstrumentationAttributes(attribute.String("111", "111")))
		//spanCtx, span := tracer.Start(context.TODO(), "main", trace.WithAttributes(attribute.String("222", "222")))
		//defer span.End()
		//
		//ctx.Next()
	}, nil
}

// NewSecondaryGinMiddleware TODO: 适用于: 次级的服务
func NewSecondaryGinMiddleware() (gin.HandlerFunc, error) {
	return func(ctx *gin.Context) {

	}, nil
}
