package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/web/httpKit"
)

// GetRequestRoute 获取请求的路由.
func GetRequestRoute(ctx *gin.Context) string {
	route := ctx.FullPath()
	if strKit.IsEmpty(route) {
		route = httpKit.GetRequestRoute(ctx.Request)
	}
	return route
}
