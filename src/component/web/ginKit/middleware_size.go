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
	// bodyLimit 单位: B
	bodyLimit := limit << 20

	return func(ctx *gin.Context) {

		// (1) do nothing
		if ctx.Request.Body == nil || ctx.Request.Body == http.NoBody {
			ctx.Next()
			return
		}

		// (2) Based on content length
		if ctx.Request.ContentLength > limit {
			ctx.AbortWithStatus(http.StatusRequestEntityTooLarge)
			return
		}

		// (3) Based on content read
		ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, bodyLimit)
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
