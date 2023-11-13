package pushKit

import (
	"net/http"
)

type InnerListener struct {
	Listener
}

func (listener InnerListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
}

func (listener InnerListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
	// 加入管理
	BindId(channel, channel.GetId())
}

func (listener InnerListener) OnMessage(channel Channel, messageType int, data []byte) {
}

func (listener InnerListener) OnClose(channel Channel, reason string, bsid, user, group string) {
	// 移除管理
	UnbindId(channel)
	UnbindBsid(channel)
	UnbindUser(channel)
	UnbindGroup(channel)
}
