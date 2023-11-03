package types

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"net/http"
)

// NewListeners
/*
PS: 本方法仅供本项目使用，严禁外部调用.
*/
func NewListeners(listener Listener) (Listeners, error) {
	if err := interfaceKit.AssertNotNil(listener, "listener"); err != nil {
		return nil, err
	}

	return []Listener{inner, listener}, nil
}

type (
	Listener interface {
		OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string)

		OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel)

		// OnMessage 收到 客户端 发来的消息.
		/*
			PS: 仅适用于WebSocket连接，因为SSE连接是单工的.
		*/
		OnMessage(channel Channel, messageType int, data []byte)

		OnClose(channel Channel, closeInfo string)
	}

	Listeners []Listener
)

func (listeners Listeners) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	for _, listener := range listeners {
		listener.OnFailure(w, r, failureInfo)
	}
}

func (listeners Listeners) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
	for _, listener := range listeners {
		listener.OnHandshake(w, r, channel)
	}
}

func (listeners Listeners) OnMessage(channel Channel, messageType int, data []byte) {
	for _, listener := range listeners {
		listener.OnMessage(channel, messageType, data)
	}
}

func (listeners Listeners) OnClose(channel Channel, closeInfo string) {
	for _, listener := range listeners {
		listener.OnClose(channel, closeInfo)
	}
}
