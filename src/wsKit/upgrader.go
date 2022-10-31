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
		/* 设置缓冲区的尺寸，以 byte 为单位 */
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,

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

// getUpgrader
/*
@return 必定不为nil
*/
func getUpgrader() *websocket.Upgrader {
	if upgrader == nil {
		return DefaultUpgrader
	}
	return upgrader
}
