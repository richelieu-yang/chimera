package wsKit

import (
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	"gitee.com/richelieu042/go-scales/src/http/httpKit"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 允许跨域
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocketHandler
// websocket服务器.
func WebSocketHandler(ctx *gin.Context) {
	// 清空cookie，防止因为cookie过大导致t-io报异常
	httpKit.ClearCookies(ctx.Request)

	// 输出所有请求头信息
	logrus.Debugf("----------------------------------------------------")
	//err := PolyfillWsConnection(ctx.Request)
	//if err != nil {
	//	logrus.Errorf("Fail to execute PolyfillWsConnection(ctx.Request), error: [%s].", err.Error())
	//}

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
			logrus.Errorf("Fail to execute ws.ReadMessage(), error: [%s].", err.Error())
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			logrus.Errorf("Fail to execute ws.WriteMessage(mt, message), error: [%s].", err.Error())
			break
		}
	}
}

// PolyfillWsConnection
// 检查并处理websocket连接.（不再强制需要nginx设置websocket穿透）
// PS: 还有问题的话，可以参考下："github.com/gorilla/websocket"中server.go的Upgrade方法.
func PolyfillWsConnection(req *http.Request) error {
	// (1)method
	if req.Method != "GET" {
		return errorKit.Simple("method(%s) isn't GET", req.Method)
	}

	// (2)proto
	// case 0: 不经过Nginx代理的话，为"HTTP/1.1"
	// case 1: 经过Nginx代理且不额外配置的话，为"HTTP/1.0"
	proto := httpKit.GetProto(req)
	switch proto {
	case "HTTP/1.0":
		fallthrough
	case "HTTP/1.1":
	default:
		return errorKit.Simple("proto(%s) is invalid", proto)
	}

	// (3)connection && upgrade
	// case 0: [http]请求，Connection: ["keep-alive"]，Upgrade: [""]
	// case 1: [ws]请求，Connection: ["Upgrade"]，Upgrade: ["websocket"]
	// case 2: nginx代理（未设置websocket穿透），[ws]请求，Connection: ["close"]，Upgrade: [""]
	// case 3: nginx代理（已设置websocket穿透），[ws]请求，Connection: ["Upgrade"]，Upgrade: ["websocket"]（Connection也有可能为"upgrade"，由于nginx配置的原因）
	connection := httpKit.GetConnection(req)
	upgrade := httpKit.GetUpgrade(req)
	if connection == "keep-alive" && upgrade == "" {
		return errorKit.Simple("req(Connection: %s, Upgrade: %s) isn't a ws request", connection, upgrade)
	}
	// 不再需要nginx设置websocket穿透了
	if connection != "Upgrade" {
		httpKit.SetConnection(req, "Upgrade")
	}
	if upgrade != "websocket" {
		httpKit.SetUpgrade(req, "websocket")
	}

	return nil
}
