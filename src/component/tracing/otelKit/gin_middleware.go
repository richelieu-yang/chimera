package otelKit

import "github.com/gin-gonic/gin"

// NewGinMiddleware
/*
PS: 需要先 SetUp || MustSetUp !!!
*/
func NewGinMiddleware() (gin.HandlerFunc, error) {
	return func(ctx *gin.Context) {
		// TODO:

	}, nil
}
