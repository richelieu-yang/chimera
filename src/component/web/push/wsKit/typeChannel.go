package wsKit

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
)

type WsChannel struct {
	*pushKit.BaseChannel

	conn    *websocket.Conn
	msgType messageType
}

func (channel *WsChannel) Push(data []byte) error {
	return channel.PushMessage(channel.msgType, data)
}

// PushMessage 推送消息给客户端.
/*
@param messageType MessageTypeText || MessageTypeBinary
*/
func (channel *WsChannel) PushMessage(messageType messageType, data []byte) (err error) {
	switch messageType {
	case MessageTypeText:
	case MessageTypeBinary:
	default:
		err = errorKit.New("invalid message type(%d)", messageType)
		return
	}
	if channel.Closed {
		return pushKit.ChannelClosedError
	}

	// 是否推送失败？
	failFlag := false

	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			err = pushKit.ChannelClosedError
			return
		}

		err = channel.conn.WriteMessage(int(messageType), data)
		failFlag = err != nil
	})

	if failFlag {
		// 推送消息失败，基本上就是连接断开了
		if channel.SetClosed() {
			info := fmt.Sprintf("Fail to push because of error(%s)", err.Error())
			channel.Listeners.OnClose(channel, info)
		}
	}

	return
}

// Close 后端主动关闭通道.
func (channel *WsChannel) Close() (err error) {
	if channel.SetClosed() {
		info := "Closed by backend"
		channel.Listeners.OnClose(channel, info)

		err = channel.conn.Close()
	}
	return
}
