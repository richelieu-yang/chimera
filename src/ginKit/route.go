// Package ginKit
// 路由相关
package ginKit

import (
	"gitee.com/richelieu042/go-scales/src/core/file/fileKit"
	"gitee.com/richelieu042/go-scales/src/core/pathKit"
	"gitee.com/richelieu042/go-scales/src/resources"
	"github.com/gin-gonic/gin"
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

	if len(methods) > 0 {
		// 接收指定类型method的请求
		for _, httpMethod := range methods {
			group.Handle(httpMethod, route, handlers...)
		}
		return
	}
	// 接收所有类型method的请求
	group.Any(route, handlers...)
}

// RegisterHandlersToRoutes 将多个相同的处理器，注册到多个路由
func RegisterHandlersToRoutes(group IGroup, routes []string, methods []string, handlers ...gin.HandlerFunc) {
	if len(handlers) == 0 {
		return
	}

	for _, relativePath := range routes {
		RegisterHandlersToRoute(group, relativePath, methods, handlers...)
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
	dir, err := pathKit.GetTempDirOfGoScales()
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
		ctx.HTML(http.StatusNotFound, fileKit.GetName(htmlPath), gin.H{
			"urlPath": ctx.Request.URL.Path,
		})
	})

	return nil
}
