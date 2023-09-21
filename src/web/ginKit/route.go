// Package ginKit
// 路由相关
package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/internal/resources"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/web/httpKit"
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

func AttachDefaultFavicon(group IGroup) error {
	iconData, err := resources.Asset("_resources/icon/favicon.ico")
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

// AttachDefaultNoRoute 使用自带的404页面.
/*
PS: 会在临时目录下生成文件，注意不要删掉他们!!!
*/
func AttachDefaultNoRoute(engine IEngine) error {
	/* 将内置的404页面解压到 临时目录 中 */
	relativePath := "_resources/html/404.html"
	tempDir, err := pathKit.GetTempDir()
	if err != nil {
		return err
	}
	err = resources.RestoreAsset(tempDir, relativePath)
	if err != nil {
		return err
	}

	/* 加载解压出的html页面 */
	htmlPath := pathKit.Join(tempDir, relativePath)
	engine.LoadHTMLFiles(htmlPath)

	NoRoute(engine, func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, fileKit.GetFileName(htmlPath), gin.H{
			"route": httpKit.GetRoute(ctx.Request),
		})
	})

	return nil
}
