package wsKit

import "github.com/gorilla/websocket"

type messageType uint

const (
	MessageTypeText   messageType = websocket.TextMessage
	MessageTypeBinary messageType = websocket.BinaryMessage
)
