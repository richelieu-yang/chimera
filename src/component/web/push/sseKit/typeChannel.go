package sseKit

import (
	"encoding/base64"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
	"net/http"
)

type SseChannel struct {
	*pushKit.BaseChannel

	w       http.ResponseWriter
	r       *http.Request
	msgType messageType
}

func (channel *SseChannel) Push(data []byte) error {
	return channel.PushMessage(channel.msgType, data)
}

// PushMessage 推送消息给客户端.
func (channel *SseChannel) PushMessage(msgType messageType, data []byte) error {
	//TODO implement me

	var str string
	switch msgType {
	case MessageTypeEncode:
		str = string(data)
		str = urlKit.EncodeURIComponent(str)
	case MessageTypeBase64:
		str = base64Kit.EncodeToString(data, base64Kit.WithEncoding(base64.StdEncoding))
	case MessageTypeRaw:
		fallthrough
	default:
		str = string(data)
	}
	event := &Event{
		Data: str,
	}

	return event.Push(channel.w)

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
	//_ = channel.SetClosed()

	if channel.SetClosed() {
		closeInfo := "Closed by backend"
		channel.Listeners.OnClose(channel, closeInfo)

		err = channel.conn.Close()
	}
	return
}
