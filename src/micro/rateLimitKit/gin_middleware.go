package rateLimitKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"golang.org/x/time/rate"
	"net/http"
)

// NewGinMiddleware Gin的限流器，用于限流.
/*
PS: 传参说明详见 NewLimiter.

@param limit			每秒生成的令牌数，	(1) 可以直接是 某个数值;
										(2) 也可以是 rate.Every() 的返回值;
@param burst 			令牌桶的容量
@param forbiddenText 	被限流时，响应给前端的内容（状态码固定为 403）
*/
func NewGinMiddleware(limit rate.Limit, burst int, forbiddenText string) gin.HandlerFunc {
	limiter := NewLimiter(limit, burst)

	if strKit.IsEmpty(forbiddenText) {
		forbiddenText = fmt.Sprintf("Exceed rate limit(%f, %d).", limit, burst)
	}

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.String(http.StatusForbidden, forbiddenText)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
