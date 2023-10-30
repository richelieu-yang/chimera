package sseKit

type MessageType uint

const (
	// MessageTypeRaw
	MessageTypeRaw MessageType = iota + 1

	// MessageTypeEncode
	MessageTypeEncode

	// MessageTypeBase64
	MessageTypeBase64
)
