package wsKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/http/httpKit"
	"github.com/sirupsen/logrus"
)

// WebSocketHandler
// websocket服务器.
func WebSocketHandler(ctx *gin.Context) {
	// 清空cookie，防止因为cookie过大导致t-io报异常
	httpKit.ClearCookies(ctx.Request)

	// 输出所有请求头信息
	logrus.Debugf("----------------------------------------------------")
	err := CheckAndPolyfill(ctx.Request)
	if err != nil {
		logrus.Errorf("Fail to execute CheckAndPolyfill(ctx.Request), error: [%s].", err.Error())
	}

	var headers map[string][]string = ctx.Request.Header
	if len(headers) != 0 {
		for k, v := range headers {
			logrus.Debugf("header(key: [%s], value: [%s]).", k, strKit.Join(v, "_"))
		}
	} else {
		logrus.Debugf("headers is empty.")
	}
	logrus.Debugf("====================================================")

	// 升级get请求为webSocket协议
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Errorf("Fail to upgrade, error: [%s].", err.Error())
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			logrus.Errorf("Fail to execute conn.ReadMessage(), error: [%s].", err.Error())
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			logrus.Errorf("Fail to execute conn.WriteMessage(mt, message), error: [%s].", err.Error())
			break
		}
	}
}
