package pushKit

type (
	MessageType uint8
)

const (
	// MessageTypeText 文本消息
	/*
		PS: 对应 gorilla/websocket 中的 websocket.TextMessage.
	*/
	MessageTypeText MessageType = iota

	// MessageTypeBinary 二进制消息
	/*
		PS: 对应 gorilla/websocket 中的: websocket.BinaryMessage.
	*/
	MessageTypeBinary
)
