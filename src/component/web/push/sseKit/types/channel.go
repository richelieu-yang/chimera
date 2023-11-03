package types

import (
	"encoding/base64"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	pushTypes "github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
	"net/http"
)

type SseChannel struct {
	*pushTypes.BaseChannel

	w       http.ResponseWriter
	r       *http.Request
	msgType MessageType
	closeCh chan string
}

// Push （写锁）推送消息给客户端.
func (channel *SseChannel) Push(data []byte) error {
	return channel.PushMessage(channel.msgType, data)
}

// PushMessage （写锁）推送消息给客户端.
func (channel *SseChannel) PushMessage(msgType MessageType, data []byte) (err error) {
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

	if channel.Closed {
		return pushKit.ChannelClosedError
	}
	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			err = pushKit.ChannelClosedError
			return
		}
		err = event.Push(channel.w)
	})
	return err
}

// Close （写锁）后端主动关闭通道.
func (channel *SseChannel) Close(reason string) error {
	if channel.SetClosed() {
		channel.closeCh <- reason
	}
	return nil
}

func (channel *SseChannel) BindGroup(group string) {
	pushKit.BindGroup(channel, group)
}

func (channel *SseChannel) BindUser(user string) {
	pushKit.BindUser(channel, user)
}

func (channel *SseChannel) BindBsid(bsid string) {
	pushKit.BindUser(channel, bsid)
}
