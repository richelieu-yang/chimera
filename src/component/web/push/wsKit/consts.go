package wsKit

import "github.com/gorilla/websocket"

type messageType struct {
	value uint
}

var (
	MessageTypeText = messageType{
		value: websocket.TextMessage,
	}
	MessageTypeBinary = messageType{
		value: websocket.BinaryMessage,
	}
)
