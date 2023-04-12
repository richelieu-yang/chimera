package ioKit

import (
	"io"
)

// NopReadCloser
/*
PS: 返回值调用Close()将什么都不会做，直接返回nil.
*/
func NopReadCloser(reader io.Reader) io.ReadCloser {
	return io.NopCloser(reader)
}

type nopWriteCloser struct {
	io.Writer
}

func (nopWriteCloser) Close() error {
	return nil
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

type nopReadSeekCloser struct {
	io.ReadSeeker
}

func (nopReadSeekCloser) Close() error {
	return nil
}

func NopReadSeekCloser(readSeeker io.ReadSeeker) io.ReadSeekCloser {
	return &nopReadSeekCloser{
		ReadSeeker: readSeeker,
	}
}
