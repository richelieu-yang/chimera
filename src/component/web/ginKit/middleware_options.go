package ginKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewOptionsMiddleware
/*
PS: 就算 NoRoute || NoMethod，也会走到中间件.
*/
func NewOptionsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodOptions {

			ctx.Header("Access-Control-Allow-Credentials", "true")
			// 预检请求的返回结果能缓存多久？24h
			ctx.Header("Access-Control-Max-Age", "86400")

			ctx.Header("Content-Type", "text/plain; charset=utf-8")
			ctx.Header("Content-Length", "0")

			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}
