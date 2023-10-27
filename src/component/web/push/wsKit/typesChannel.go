package wsKit

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
)

type WsChannel struct {
	pushKit.Channel
	pushKit.BaseChannel

	conn *websocket.Conn
}

func (channel *WsChannel) Push(messageType pushKit.MessageType, data []byte) (err error) {
	if channel.Closed {
		return pushKit.ChannelClosedError
	}

	// 写锁
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			err = pushKit.ChannelClosedError
			return
		}

		var t int
		switch messageType {
		case pushKit.MessageTypeBinary:
			t = websocket.BinaryMessage
		case pushKit.MessageTypeText:
			fallthrough
		default:
			t = websocket.TextMessage
		}
		err = channel.conn.WriteMessage(t, data)
	})

	if err != nil && !errors.Is(err, pushKit.ChannelClosedError) {
		// 推送消息失败，基本上就是连接断开了
		if channel.SetClosed() {
			info := fmt.Sprintf("Fail to push because of error(%s)", err.Error())
			channel.Listener.OnClose(channel, info)
		}
	}

	return
}

func (channel *WsChannel) Close() (err error) {
	if channel.SetClosed() {
		info := "Closed by backend"
		channel.Listener.OnClose(channel, info)

		err = channel.conn.Close()
	}
	return
}
