package sseKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"net/http"
)

type SseChannel struct {
	*pushKit.BaseChannel

	w       http.ResponseWriter
	r       *http.Request
	msgType messageType
}

func (channel *SseChannel) Push(data []byte) error {
	//TODO implement me
	panic("implement me")
}

// PushMessage 推送消息给客户端.
func (channel *SseChannel) PushMessage(t messageType, data []byte) (err error) {
	//if channel.Closed {
	//	return pushKit.ChannelClosedError
	//}
	//
	//// 写锁
	//channel.RWMutex.LockFunc(func() {
	//	if channel.Closed {
	//		err = pushKit.ChannelClosedError
	//		return
	//	}
	//
	//	_, err := fmt.Fprintf(channel.w, "%s", msg)
	//	if err != nil {
	//		logrus.WithError(err).Error("fail to send initial message")
	//		return
	//	}
	//	channel.w.(http.Flusher).Flush()
	//
	//})

	panic("TODO")

	return nil
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
