package wsKit

import (
	"compress/flate"
	"github.com/gorilla/websocket"
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

// newChannel
/*
@param conn 必定不为nil
*/
func newChannel(conn *websocket.Conn) (*Channel, error) {
	c := &Channel{
		conn: conn,
		lock: new(sync.Mutex),
	}

	// 写（推送）的压缩 compress
	conn.EnableWriteCompression(true)
	if err := conn.SetCompressionLevel(flate.BestSpeed); err != nil {
		return nil, err
	}

	if listener != nil {
		listener.OnAfterHandshake(c)

		conn.SetCloseHandler(func(code int, text string) error {
			listener.OnAfterClose(c, code, text)
			return nil
		})

		// 接收前端发来的数据
		go func(c *Channel) {
			var err error

			for {
				msgType, msgData, err := conn.ReadMessage()
				if err != nil {
					// 一旦读取失败，就中断读循环
					break
				}
				listener.OnMessage(c, msgType, msgData)
			}
			getLogger().Warnf("[READ] Loop stops because of error: %+v", err)
		}(c)
	} else {
		getLogger().Warn("[READ] Won't process data from front end because listener == nil.")
	}

	return c, nil
}
