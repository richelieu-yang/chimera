package ioKit

import (
	"io"
)

// NopCloser
/*
PS: 返回值调用Close()将什么都不会做，直接返回nil.
*/
var NopCloser func(reader io.Reader) io.ReadCloser = io.NopCloser

// NopCloserToReadSeeker 给 io.ReadSeeker 实例加上Close()（虽然这个Close()什么都不做）.
func NopCloserToReadSeeker(readSeeker io.ReadSeeker) io.ReadSeekCloser {
	return &readSeekCloser{
		readSeeker: readSeeker,
	}
}
