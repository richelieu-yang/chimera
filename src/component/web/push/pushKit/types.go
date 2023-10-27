package pushKit

import "net/http"

type Listener interface {
	OnFailure(w http.ResponseWriter, r *http.Request, error string)

	OnHandshake(w http.ResponseWriter, r *http.Request)

	// OnMessage 收到 客户端 发来的消息.
	/*
		PS: 仅适用于WebSocket连接，因为SSE连接是单工的.
	*/
	OnMessage(channel Channel, messageType MessageType, data []byte)

	OnClose(channel Channel, code int, text string)
}

type Channel interface {
	// Push 推送（二进制）消息给客户端.
	Push(messageType MessageType, data []byte) error

	// Close 后端主动关闭通道.
	Close() error
}
