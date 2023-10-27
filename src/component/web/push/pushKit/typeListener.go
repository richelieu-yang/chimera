package pushKit

import "net/http"

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
