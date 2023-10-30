package wsKit

import "github.com/gorilla/websocket"

type MessageType uint

const (
	MessageTypeText   MessageType = websocket.TextMessage
	MessageTypeBinary MessageType = websocket.BinaryMessage
)
