// Package ginKit
// 路由相关
package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/resources"
	"net/http"
)

// RegisterHandlersToRoute 将多个相同的处理器，注册到1个路由
/*
@param methods nil => 接收所有类型method的请求.	e.g. http.MethodGet、http.MethodPost
*/
func RegisterHandlersToRoute(group IGroup, route string, methods []string, handlers ...gin.HandlerFunc) {
	if len(handlers) == 0 {
		return
	}

	// (1) Any
	if len(methods) == 0 {
		group.Any(route, handlers...)
	}
	// (2) 指定类型的method
	for _, method := range methods {
		group.Handle(method, route, handlers...)
	}
	return
}

// RegisterHandlersToRoutes 将多个相同的处理器，注册到多个路由
func RegisterHandlersToRoutes(group IGroup, routes []string, methods []string, handlers ...gin.HandlerFunc) {
	for _, route := range routes {
		RegisterHandlersToRoute(group, route, methods, handlers...)
	}
}

func FaviconByDefault(group IGroup) error {
	iconData, err := resources.Asset("resources/icon/favicon.ico")
	if err != nil {
		return err
	}

	group.Any("/favicon.ico", func(ctx *gin.Context) {
		RespondIcon(ctx, http.StatusOK, iconData)
	})
	return nil
}

// NoRoute 404
func NoRoute(engine IEngine, handlers ...gin.HandlerFunc) {
	engine.NoRoute(handlers...)
}

// NoRouteByDefault 使用自带的404页面
func NoRouteByDefault(engine IEngine) error {
	relPath := "resources/html/404.html"

	/* 将内置的404页面解压到临时目录中 */
	dir, err := pathKit.GetUniqueTempDir()
	if err != nil {
		return err
	}
	err = resources.RestoreAsset(dir, relPath)
	if err != nil {
		return err
	}

	htmlPath := pathKit.Join(dir, relPath)
	LoadHtmlFiles(engine, htmlPath)
	NoRoute(engine, func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, fileKit.GetBaseName(htmlPath), gin.H{
			"urlPath": ctx.Request.URL.Path,
		})
	})

	return nil
}
