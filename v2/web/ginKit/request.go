package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/core/strKit"
	"github.com/richelieu42/chimera/v2/web/httpKit"
)

// GetRequestRoute 获取请求的路由.
func GetRequestRoute(ctx *gin.Context) string {
	route := ctx.FullPath()
	if strKit.IsEmpty(route) {
		route = httpKit.GetRequestRoute(ctx.Request)
	}
	return route
}
