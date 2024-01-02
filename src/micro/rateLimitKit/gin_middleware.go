package rateLimitKit

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

// NewGinMiddleware Gin的限流器，用于限流.
/*
PS: 传参说明详见 NewLimiter.
*/
func NewGinMiddleware(limit rate.Limit, burst int) gin.HandlerFunc {
	limiter := NewLimiter(limit, burst)

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.String(http.StatusForbidden, "Exceed rate limit.")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
