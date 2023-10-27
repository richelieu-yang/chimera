package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
	"net/http"
)

type Listener interface {
	OnFailure(w http.ResponseWriter, r *http.Request, error string)

	OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel)

	// OnMessage 收到 客户端 发来的消息.
	/*
		PS: 仅适用于WebSocket连接，因为SSE连接是单工的.
	*/
	OnMessage(channel Channel, messageType int, data []byte)

	OnClose(channel Channel, code int, text string)
}

type Channel interface {
	// Push 推送（二进制）消息给客户端.
	Push(messageType MessageType, data []byte) error

	// Close 后端主动关闭通道.
	Close() error
}

type BaseChannel struct {
	Id    string
	Bsid  string
	User  string
	Group string

	RWMutex mutexKit.RWMutex

	Data   interface{}
	Closed bool
}

func (channel *BaseChannel) IsClosed() (rst bool) {
	channel.RWMutex.RLockFunc(func() {
		rst = channel.Closed
	})
	return
}

func (channel *BaseChannel) SetClosed() {
	if channel.Closed {
		return
	}

	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			return
		}
		channel.Closed = true
	})
}
