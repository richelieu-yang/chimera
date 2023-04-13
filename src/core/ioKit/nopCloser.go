package ioKit

import (
	"io"
)

type nopCloserToWriter struct {
	io.Writer
}

func (nopCloserToWriter) Close() error {
	return nil
}

type nopCloserToReadSeeker struct {
	io.ReadSeeker
}

func (nopCloserToReadSeeker) Close() error {
	return nil
}

// NopCloserToReader
/*
PS: 返回值调用Close()将什么都不会做，直接返回nil.
*/
func NopCloserToReader(reader io.Reader) io.ReadCloser {
	return io.NopCloser(reader)
}

// NopCloserToWriter
/*
PS: 返回值调用Close()将什么都不会做，直接返回nil.
*/
func NopCloserToWriter(writer io.Writer) io.WriteCloser {
	return &nopCloserToWriter{
		writer,
	}
}

func NopCloserToReadSeeker(readSeeker io.ReadSeeker) io.ReadSeekCloser {
	return &nopCloserToReadSeeker{
		ReadSeeker: readSeeker,
	}
}
