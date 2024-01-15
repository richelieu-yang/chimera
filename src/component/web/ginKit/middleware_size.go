package ginKit

import (
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewSizeLimiterMiddleware 参考了echo中的 middleware.BodyLimit()
/*
	@param limit 单位: MiB
*/
func NewSizeLimiterMiddleware(limit int64) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request

		// (1) do nothing
		if req.Body == nil || req.Body == http.NoBody {
			ctx.Next()
			return
		}

		// (2) Based on content length
		if req.ContentLength > limit {
			ctx.AbortWithStatus(http.StatusRequestEntityTooLarge)
			return
		}

		// (3) Based on content read
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, limit)
		ctx.Next()
	}
}

// NewSizeLimiterMiddleware1
/*
gin-contrib/size
	https://github.com/gin-contrib/size
*/
func NewSizeLimiterMiddleware1(limit int64) gin.HandlerFunc {
	return limits.RequestSizeLimiter(limit)
}
