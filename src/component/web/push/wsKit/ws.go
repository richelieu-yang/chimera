package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"net/http"
	"time"
)

// DefaultUpgrader 默认的Upgrader.
/*
@return 并发安全的
*/
func DefaultUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		HandshakeTimeout: time.Second * 3,
		CheckOrigin: func(r *http.Request) bool {
			// 允许跨域
			return true
		},
	}
}

// NewProcessor
/*
@param upgrader			可以为nil
@param idGenerator		可以为nil（使用xid）
@param listener			不能为nil
*/
func NewProcessor(upgrader *websocket.Upgrader, idGenerator func() (string, error), listener pushKit.Listener, messageType messageType) (*WsProcessor, error) {
	if upgrader == nil {
		upgrader = DefaultUpgrader()
	}
	if idGenerator == nil {
		idGenerator = func() (string, error) {
			return idKit.NewXid(), nil
		}
	}
	if err := interfaceKit.AssertNotNil(listener, "listener"); err != nil {
		return nil, err
	}
	switch messageType {
	case MessageTypeText:
	case MessageTypeBinary:
	default:
		return nil, errorKit.New("invalid message type(%d)", messageType)
	}

	return &WsProcessor{
		upgrader:           upgrader,
		idGenerator:        idGenerator,
		listener:           listener,
		defaultMessageType: messageType,
	}, nil
}
