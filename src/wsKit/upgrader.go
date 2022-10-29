package wsKit

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	// 真正使用的Upgrader
	upgrader *websocket.Upgrader = nil

	DefaultUpgrader = &websocket.Upgrader{
		// 握手的超时时间
		HandshakeTimeout: time.Second * 6,

		CheckOrigin: func(r *http.Request) bool {
			// 允许跨域
			return true
		},
	}
)

func setUpgrader(upgrader1 *websocket.Upgrader) {
	if upgrader1 != nil {
		upgrader = upgrader1
	}
}

func getUpgrader() *websocket.Upgrader {
	if upgrader == nil {
		return DefaultUpgrader
	}
	return upgrader
}
