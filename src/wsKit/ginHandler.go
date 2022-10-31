package wsKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/http/httpKit"
)

// GinHandler
/*
PS:
(1) 调用此方法前应该调用 wsKit.Initialize()；
(2) 能接受Any类型的请求，但非GET类型的握手不会成功.
*/
func GinHandler(ctx *gin.Context) {
	// 清空cookie，防止因为cookie过大导致t-io报异常
	httpKit.ClearCookies(ctx.Request)

	conn, err := getUpgrader().Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	if err != nil {
		getLogger().Errorf("fail to upgrade, error: %+v", err)
		return
	}

	_, err = newChannel(conn)
	if err != nil {
		getLogger().Errorf("fail to new channel, error: %+v", err)
		return
	}
}
