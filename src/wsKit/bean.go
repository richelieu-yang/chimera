package wsKit

import "github.com/gorilla/websocket"

type (
	wsConnection struct {
		// gorilla/websocket的连接
		ws *websocket.Conn
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

func (conn *wsConnection) Dispose() {
	if conn == nil {
		return
	}
	_ = conn.ws.Close()
	conn.ws = nil
}

func NewWsConnection() *wsConnection {
	return new(wsConnection)
}
