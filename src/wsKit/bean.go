package wsKit

import (
	"github.com/gorilla/websocket"
	"sync"
)

type (
	Connection struct {
		// 唯一id
		uniqueId string
		// 锁
		lock *sync.Mutex

		// gorilla/websocket的连接
		conn *websocket.Conn

		// 所属于的群组（有且仅有一个）
		group string
		// 所属于的用户（有且仅有一个）
		user  string
		token string
		data  map[string]interface{}
	}
)

func (c *Connection) Dispose() error {
	conn := c.conn
	if conn == nil {
		return NoConnError
	}

	err := c.conn.Close()
	c.conn = nil
	return err
}

func (c *Connection) Push(messageType int, data []byte) error {

	ws := c.conn
	if err := ws.WriteMessage(messageType, data); err != nil {
		return err
	}
	return nil
}

func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		conn: conn,
		lock: new(sync.Mutex),
	}
}
