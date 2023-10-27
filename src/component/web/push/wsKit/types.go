package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"net/http"
	"time"
)

type (
	Handler struct {
		upgrader *websocket.Upgrader

		idGenerator func() (string, error)
	}
)

func NewHandler(handshakeTimeout time.Duration, checkOrigin func(r *http.Request) bool, idGenerator func() (string, error)) {
	if checkOrigin == nil {
		checkOrigin = func(r *http.Request) bool {
			// 允许跨域
			return true
		}
	}
	if idGenerator == nil {
		idGenerator = func() (string, error) {
			return idKit.NewXid(), nil
		}
	}

	// upgrader 并发安全的
	var upgrader = websocket.Upgrader{
		HandshakeTimeout: time.Second * 3,
		CheckOrigin: func(r *http.Request) bool {
			// 允许跨域
			return true
		},
	}
}
