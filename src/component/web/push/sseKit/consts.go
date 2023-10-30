package sseKit

type messageType uint

const (
	// MessageTypeRaw
	MessageTypeRaw messageType = iota + 1

	// MessageTypeEncode
	MessageTypeEncode

	// MessageTypeBase64
	MessageTypeBase64
)
