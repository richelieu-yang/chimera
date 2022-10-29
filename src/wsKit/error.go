package wsKit

import "github.com/gorilla/websocket"

// IsCloseError 判断是否是 CloseError
/*
e.g.
向已经关闭的Conn读取数据，会返回 CloseError.
*/
func IsCloseError(err error) (*websocket.CloseError, bool) {
	ce, ok := err.(*websocket.CloseError)
	return ce, ok
}

// IsExpectedCloseError 错误判断
func IsExpectedCloseError(err error, codes ...int) bool {
	return websocket.IsCloseError(err, codes...)
}

// IsUnexpectedCloseError 错误判断
func IsUnexpectedCloseError(err error, expectedCodes ...int) bool {
	return websocket.IsUnexpectedCloseError(err, expectedCodes...)
}
