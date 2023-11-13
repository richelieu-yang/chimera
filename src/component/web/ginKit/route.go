// Package ginKit
// 路由相关
package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/internal/resources"
	"net/http"
)

// RegisterRoute
/*
适用场景: 1个路由，n个Method.

@param methods nil => 接收所有类型method的请求.	e.g. http.MethodGet、http.MethodPost
*/
func RegisterRoute(group IGroup, route string, methods []string, handlers ...gin.HandlerFunc) {
	if len(handlers) == 0 {
		return
	}

	if len(methods) == 0 {
		// (1) Any
		group.Any(route, handlers...)
	} else {
		// (2) 指定类型的method
		for _, method := range methods {
			group.Handle(method, route, handlers...)
		}
	}
}

// RegisterRoutes 将多个相同的处理器，注册到多个路由.
func RegisterRoutes(group IGroup, routes []string, methods []string, handlers ...gin.HandlerFunc) {
	for _, route := range routes {
		RegisterRoute(group, route, methods, handlers...)
	}
}

func AttachDefaultFavicon(engine *gin.Engine) error {
	iconData, err := resources.Asset("_resources/icon/favicon.ico")
	if err != nil {
		return err
	}

	engine.GET("/favicon.ico", func(ctx *gin.Context) {
		RespondIcon(ctx, http.StatusOK, iconData)
	})
	return nil
}
