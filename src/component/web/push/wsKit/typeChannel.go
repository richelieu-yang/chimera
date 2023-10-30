package wsKit

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
)

type WsChannel struct {
	*pushKit.BaseChannel

	conn *websocket.Conn
}

// Push 推送消息给客户端.
/*
@param messageType websocket.TextMessage || websocket.BinaryMessage
*/
func (channel *WsChannel) Push(messageType int, data []byte) (err error) {
	if channel.Closed {
		return pushKit.ChannelClosedError
	}

	// 是否推送失败？
	flag := false

	// 写锁
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			err = pushKit.ChannelClosedError
			return
		}

		switch messageType {
		case websocket.TextMessage:
		case websocket.BinaryMessage:
		default:
			err = errorKit.New("invalid WebSocket message type(%d)", messageType)
			return
		}
		err = channel.conn.WriteMessage(messageType, data)
		flag = err != nil
	})

	if flag {
		// 推送消息失败，基本上就是连接断开了
		if channel.SetClosed() {
			info := fmt.Sprintf("Fail to push because of error(%s)", err.Error())
			channel.Listener.OnClose(channel, info)
		}
	}

	return
}

// Close 后端主动关闭通道.
func (channel *WsChannel) Close() (err error) {
	if channel.SetClosed() {
		info := "Closed by backend"
		channel.Listener.OnClose(channel, info)

		err = channel.conn.Close()
	}
	return
}
