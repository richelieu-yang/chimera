package ioKit

import (
	"io"
)

type nopWriteCloser struct {
	io.Writer
}

func (nopWriteCloser) Close() error {
	return nil
}

// NopReadCloser
/*
PS: 返回值调用Close()将什么都不会做，直接返回nil.
*/
func NopReadCloser(reader io.Reader) io.ReadCloser {
	return io.NopCloser(reader)
}

// NopWriteCloser
/*
PS: 返回值调用Close()将什么都不会做，直接返回nil.
*/
func NopWriteCloser(writer io.Writer) io.WriteCloser {
	return &nopWriteCloser{
		writer,
	}
}
