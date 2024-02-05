package otelKit

import "github.com/gin-gonic/gin"

// NewGinMiddleware
/*
PS: 需要先setup!!!
*/
func NewGinMiddleware() (gin.HandlerFunc, error) {
	if err := check(); err != nil {
		return nil, err
	}

	return func(ctx *gin.Context) {

	}, nil
}
