package rateLimitKit

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

// NewRateLimitMiddleware Gin的限流器，用于限流.
func NewRateLimitMiddleware(r rate.Limit, b int) gin.HandlerFunc {
	limiter := NewLimiter(r, b)

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.String(http.StatusForbidden, "Exceed rate limit.")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
