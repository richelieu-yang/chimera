package wsKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/http/httpKit"
)

func WebSocketHandler(ctx *gin.Context) {
	// 清空cookie，防止因为cookie过大导致t-io报异常
	httpKit.ClearCookies(ctx.Request)

}
