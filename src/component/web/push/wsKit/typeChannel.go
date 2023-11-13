package wsKit

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
)

type WsChannel struct {
	pushKit.BaseChannel

	conn        *websocket.Conn
	messageType messageType
}

func (channel *WsChannel) Push(data []byte) error {
	return channel.PushMessage(channel.messageType, data)
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

		err = channel.conn.WriteMessage(messageType.value, data)
		failFlag = err != nil
	})

	if failFlag {
		// 推送消息失败，基本上就是连接断开了
		closeInfo := fmt.Sprintf("Fail to push because of error(%s)", err.Error())
		if channel.SetClosed() {
			channel.CloseCh <- closeInfo
		}
	}

	return
}

// Close 后端主动关闭通道.
func (channel *WsChannel) Close(reason string) (err error) {
	closeInfo := fmt.Sprintf("Closed by backend with reason(%s)", reason)
	if channel.SetClosed() {
		channel.CloseCh <- closeInfo
		err = channel.conn.Close()
	}
	return
}

func (channel *WsChannel) BindGroup(group string) {
	pushKit.BindGroup(channel, group)
}

func (channel *WsChannel) BindUser(user string) {
	pushKit.BindUser(channel, user)
}

func (channel *WsChannel) BindBsid(bsid string) {
	pushKit.BindBsid(channel, bsid)
}
