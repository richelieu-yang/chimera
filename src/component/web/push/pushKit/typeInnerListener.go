package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/bytesKit"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
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

	// 仅针对SSE连接
	if listener.sseFlag {
		var err error
		defer func() {
			if err != nil {
				listener.c = nil
			}
		}()

		listener.c = cronKit.NewCron()
		_, err = listener.c.AddFunc("@every 15s", func() {
			if err := channel.Push(pongData); err != nil {
				logger.WithError(err).Error("Fail to pong")
				return
			}
		})
		if err != nil {
			logger.WithError(err).WithField("channelId", channel.GetId()).Error("Fail to AddFunc()")
			return
		}
		// 不会阻塞地启动
		listener.c.Start()
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
