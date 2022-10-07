package ioKit

import (
	"bytes"
)

// NewReadWriter
/*
bytes.Buffer 结构体 实现了 io.ReadWriter 接口
*/
func NewReadWriter(s []byte) *bytes.Buffer {
	return bytes.NewBuffer(s)
}

// NewReadWriterFromString
/*
bytes.Buffer 结构体 实现了 io.ReadWriter 接口
*/
func NewReadWriterFromString(str string) *bytes.Buffer {
	return bytes.NewBufferString(str)
}
