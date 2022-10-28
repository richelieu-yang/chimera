package wsKit

import (
	"compress/flate"
	"github.com/gorilla/websocket"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"sync"
)

type (
	Channel struct {
		disposed bool

		// 锁
		lock *sync.Mutex

		// gorilla/websocket的连接
		conn *websocket.Conn

		// 唯一id
		uniqueId string
		// 所属于的群组（有且仅有一个）
		group string
		// 所属于的用户（有且仅有一个）
		user  string
		token string
		data  map[string]interface{}
	}

	Listener interface {
		// 接收到前端发来的消息
		onMessage(c *Channel, msgType int, msgData []byte)

		// 监听 websocket 连接断开
		onClose(c *Channel)
	}
)

func (c *Channel) Dispose() error {
	if c == nil {
		return nil
	}

	conn := c.conn
	c.conn = nil
	//if conn == nil {
	//	return NoConnError
	//}
	return conn.Close()
}

func (c *Channel) PushTextMessage(msgData []byte) error {
	return c.PushMessage(websocket.TextMessage, msgData)
}

func (c *Channel) PushBinaryMessage(msgData []byte) error {
	return c.PushMessage(websocket.BinaryMessage, msgData)
}

func (c *Channel) PushMessage(msgType int, msgData []byte) error {
	if c == nil {
		return nil
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	ws := c.conn
	if err := ws.WriteMessage(msgType, msgData); err != nil {
		return err
	}
	return nil
}

func NewConnection(conn *websocket.Conn, listener Listener) (*Channel, error) {
	lock := new(sync.Mutex)
	if conn == nil {
		return nil, errorKit.Simple("conn == nil")
	}

	c := &Channel{
		conn: conn,
		lock: lock,
	}

	// 写（推送）的压缩 compress
	conn.EnableWriteCompression(true)
	if err := conn.SetCompressionLevel(flate.BestSpeed); err != nil {
		return nil, err
	}

	// 接收前端发来的数据
	if listener != nil {
		go func() {
			var err error
			for {
				msgType, msgData, err := conn.ReadMessage()
				if err != nil {
					break
				}
				listener.onMessage(c, msgType, msgData)
			}

			// TODO: 接收数据失败，此时默认该连接废了？
			err = err
		}()
	}

	return c, nil
}
