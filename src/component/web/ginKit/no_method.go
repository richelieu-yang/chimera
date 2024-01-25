package ginKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"net/http"
)

func DefaultNoMethod(engine *gin.Engine) {
	engine.NoMethod(func(ctx *gin.Context) {
		var allowed []string
		route := ctx.Request.URL.Path
		routeInfoSlice := engine.Routes()
		for _, routeInfo := range routeInfoSlice {
			if routeInfo.Path == route {
				allowed = append(allowed, routeInfo.Method)
				AddResponseHeader(ctx, "Allow", routeInfo.Method)
			}
		}

		text := fmt.Sprintf("Method(%s) isn't allowed for route(%s) and allowed methods is %s.", ctx.Request.Method, ctx.Request.URL.Path, allowed)
		if strKit.IsNotEmpty(serviceInfo) {
			text = fmt.Sprintf("[%s] %s", serviceInfo, text)
		}
		ctx.String(http.StatusMethodNotAllowed, text)
	})
}
