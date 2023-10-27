package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"net/http"
	"time"
)

type (
	Handler struct {
		// upgrader 是并发安全的
		upgrader *websocket.Upgrader

		idGenerator func() (string, error)
	}
)

// NewHandler
/*
@param handshakeTimeout	默认3s
@param checkOrigin		可以为nil（允许所有）
@param idGenerator		可以为nil（使用xid）
*/
func NewHandler(handshakeTimeout time.Duration, checkOrigin func(r *http.Request) bool, idGenerator func() (string, error)) *Handler {
	handshakeTimeout = timeKit.ToDefaultDurationIfInvalid(handshakeTimeout, time.Second*3)
	if checkOrigin == nil {
		checkOrigin = func(r *http.Request) bool {
			// 允许跨域
			return true
		}
	}
	// upgrader 并发安全的
	upgrader := &websocket.Upgrader{
		HandshakeTimeout: handshakeTimeout,
		CheckOrigin: func(r *http.Request) bool {
			// 允许跨域
			return true
		},
	}

	if idGenerator == nil {
		idGenerator = func() (string, error) {
			return idKit.NewXid(), nil
		}
	}

	return &Handler{
		upgrader:    upgrader,
		idGenerator: idGenerator,
	}
}
