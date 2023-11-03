package types

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"net/http"
)

type InnerListener struct {
	Listener
}

func (listener InnerListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
}

func (listener InnerListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
	// 加入管理
	pushKit.BindId(channel, channel.GetId())
}

func (listener InnerListener) OnMessage(channel Channel, messageType int, data []byte) {
}

func (listener InnerListener) OnClose(channel Channel, closeInfo string) {
	// 移除管理
	pushKit.UnBindId(channel)
	pushKit.UnbindBsid(channel)
	pushKit.UnbindUser(channel)
	pushKit.UnbindGroup(channel)
}
