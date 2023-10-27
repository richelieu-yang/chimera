package pushKit

import (
	"net/http"
	"sync"
)

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

type BaseChannel struct {
	Id    string `json:"id"`
	Bsid  string `json:"bsid"`
	User  string `json:"user"`
	Group string `json:"group"`

	Lock sync.Mutex `json:"lock"`

	Data   interface{} `json:"data"`
	Closed bool        `json:"closed"`
}

type Channel interface {
	// Push 推送（二进制）消息给客户端.
	Push(messageType MessageType, data []byte) error

	// Close 后端主动关闭通道.
	Close() error
}
