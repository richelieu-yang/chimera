package pushKit

import "net/http"

var mListener = &managerListener{}

func NewListeners(listener Listener) Listeners {
	return []Listener{mListener, listener}
}

type Listener interface {
	OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string)

	OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel)

	// OnMessage 收到 客户端 发来的消息.
	/*
		PS: 仅适用于WebSocket连接，因为SSE连接是单工的.
	*/
	OnMessage(channel Channel, messageType int, data []byte)

	OnClose(channel Channel, closeInfo string)
}

type Listeners []Listener

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

type managerListener struct {
	Listener
}

func (listener managerListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {

}

func (listener managerListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
}

func (listener managerListener) OnMessage(channel Channel, messageType int, data []byte) {
	// TODO: 加入管理（manager.go）
}

func (listener managerListener) OnClose(channel Channel, closeInfo string) {
	// TODO: 移除管理（manager.go）
}
