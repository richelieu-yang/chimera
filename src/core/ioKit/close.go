package ioKit

import (
	"io"
)

// CloseWriter 尝试关闭 传参writer（如果能关闭的话）
/*
PS: 如果 传参obj 实现了 io.Closer 接口，且非特殊值，尝试关闭它；否则直接返回nil.
*/
func CloseWriter(writer io.Writer) error {
	if closer, ok := writer.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// CloseReader
/*
@param reader 可以为nil
*/
func CloseReader(reader io.Reader) error {
	if closer, ok := reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
