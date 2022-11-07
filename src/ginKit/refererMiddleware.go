package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/http/refererKit"
	"net/http"
)

// NewRefererMiddleware 验证请求的referer的中间件
func NewRefererMiddleware(none bool, blocked bool, route string, serverNames ...string) (gin.HandlerFunc, error) {
	verifier, err := refererKit.NewRefererVerifier(none, blocked, route, serverNames...)
	if err != nil {
		return nil, err
	}
	return func(ctx *gin.Context) {
		route := ctx.FullPath()
		referer := ctx.GetHeader("referer")

		if ok, reason := verifier.Verify(route, referer); ok {
			ctx.Next()
		} else {
			// http.StatusUnauthorized: 401
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"reason": reason,
			})
		}
	}, nil
}
