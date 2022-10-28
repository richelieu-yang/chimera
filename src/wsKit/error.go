package wsKit

import "github.com/gorilla/websocket"

// IsCloseError 错误判断
func IsCloseError(err error) bool {
	return websocket.IsCloseError(err)
}

// IsUnexpectedCloseError 错误判断
func IsUnexpectedCloseError(err error) bool {
	return websocket.IsUnexpectedCloseError(err)
}
