// Package ginKit
// 路由相关
package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
)

// RegisterHandlers
/*
适用场景: 1个路由，n个Method.

@param methods 	nil => 接收所有类型method的请求.	e.g. http.MethodGet、http.MethodPost
@param handlers 其中的元素不能为nil!!!
*/
func RegisterHandlers(group IGroup, route string, methods []string, handlers ...gin.HandlerFunc) (err error) {
	if len(handlers) == 0 {
		return
	}
	sliceKit.Each(handlers, func(handler gin.HandlerFunc, index int) bool {
		if handler == nil {
			err = errorKit.New("handlers[%d] == nil", index)
			return true
		}
		return false
	})
	if err != nil {
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
	return
}

// RegisterHandlersRoutes 将相同的多个处理器，注册到多个路由.
func RegisterHandlersRoutes(group IGroup, routes []string, methods []string, handlers ...gin.HandlerFunc) (err error) {
	routes = sliceKit.Uniq(routes)

	for _, route := range routes {
		err = RegisterHandlers(group, route, methods, handlers...)
		if err != nil {
			return
		}
	}
	return
}
