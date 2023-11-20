package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/bytesKit"
	"github.com/robfig/cron/v3"
	"net/http"
)

var (
	pingData = []byte("ping")
	pongData = []byte("pong")
)

type innerListener struct {
	Listener

	sseFlag bool
	c       *cron.Cron
}

func (listener innerListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
}

func (listener innerListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
	// 加入管理
	BindId(channel, channel.GetId())

	if listener.sseFlag {

	}
}

func (listener innerListener) OnMessage(channel Channel, messageType int, data []byte) {
	// 仅针对WebSocket连接
	if bytesKit.Equals(data, pingData) {
		if err := channel.Push(pongData); err != nil {
			logger.WithError(err).Error("Fail to pong")
			return
		}
	}
}

func (listener innerListener) OnClose(channel Channel, closeInfo string, bsid, user, group string) {
	// 移除管理
	UnbindId(channel)
	UnbindBsid(channel)
	UnbindUser(channel)
	UnbindGroup(channel)
}
