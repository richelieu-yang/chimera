package types

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"net/http"
)

var inner = &innerListener{}

type innerListener struct {
	Listener
}

func (listener innerListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
}

func (listener innerListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
	// 加入管理
	pushKit.BindId(channel, channel.GetId())
}

func (listener innerListener) OnMessage(channel Channel, messageType int, data []byte) {
}

func (listener innerListener) OnClose(channel Channel, closeInfo string) {
	// 移除管理
	channel.Unbind()
}
