package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/http/httpKit"
	"net/http"
)

// SetWebSocketPenetration 设置websocket穿透
/*
PS:
(1) 当客户使用了Nginx，但不设置websocket穿透，我们只能通过Golang代码来设置websocket穿透了；
(2) 先调用此方法，再调用 IsWebSocketUpgrade().
*/
func SetWebSocketPenetration(req *http.Request) {
	/*
		case 0: [http]请求，Connection: ["keep-alive"]，Upgrade: [""]
		case 1: [conn]请求，Connection: ["Upgrade"]，Upgrade: ["websocket"]
		case 2: nginx代理（未设置websocket穿透），[conn]请求，Connection: ["close"]，Upgrade: [""]
		case 3: nginx代理（已设置websocket穿透），[conn]请求，Connection: ["Upgrade"]，Upgrade: ["websocket"]（Connection也有可能为"upgrade"，由于nginx配置的原因）
	*/
	if req.Method == http.MethodGet {
		connection := httpKit.GetHeader(req.Header, "Connection")
		upgrade := httpKit.GetHeader(req.Header, "Upgrade")

		if strKit.EqualsIgnoreCase(connection, "close") && strKit.IsEmpty(upgrade) {
			// 代码设置websocket穿透
			httpKit.SetHeader(req.Header, "Connection", "Upgrade")
			httpKit.SetHeader(req.Header, "Upgrade", "websocket")
		}
	}
}

// IsWebSocketUpgrade
/*
PS: 建议先调用 SetWebSocketPenetration().

@return true: websocket请求; false: 普通http请求.
*/
func IsWebSocketUpgrade(req *http.Request) bool {
	return websocket.IsWebSocketUpgrade(req)
}

// AssertWebSocketUpgrade
/*
PS: 建议先调用 SetWebSocketPenetration().
*/
func AssertWebSocketUpgrade(req *http.Request) error {
	if !websocket.IsWebSocketUpgrade(req) {
		connection := httpKit.GetHeader(req.Header, "Connection")
		upgrade := httpKit.GetHeader(req.Header, "Upgrade")
		return errorKit.Simple("request(connection: %s, upgrade: %s) isn't a websocket request", connection, upgrade)
	}
	return nil
}
