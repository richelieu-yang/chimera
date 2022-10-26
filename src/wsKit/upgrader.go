package wsKit

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	// 握手的超时时间
	HandshakeTimeout: time.Second * 6,

	CheckOrigin: func(r *http.Request) bool {
		// 允许跨域
		return true
	},
}
