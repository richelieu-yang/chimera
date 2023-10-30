package sseKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/wsKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SseChannel struct {
	pushKit.Channel
	pushKit.Channel

	w http.ResponseWriter
	r *http.Request
}

// Push 推送消息给客户端.
func (channel *SseChannel) Push(messageType wsKit.MessageType, data []byte) (err error) {
	if channel.Closed {
		return pushKit.ChannelClosedError
	}

	// 写锁
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			err = pushKit.ChannelClosedError
			return
		}

		_, err := fmt.Fprintf(channel.w, "%s", msg)
		if err != nil {
			logrus.WithError(err).Error("fail to send initial message")
			return
		}
		channel.w.(http.Flusher).Flush()

	})

}

// Close 后端主动关闭通道.
func (channel *SseChannel) Close() (err error) {
	_ = channel.SetClosed()

	//if channel.SetClosed() {
	//	info := "Closed by backend"
	//	channel.Listener.OnClose(channel, info)
	//
	//	err = channel.conn.Close()
	//}
	return
}
