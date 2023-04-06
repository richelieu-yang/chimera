package ioKit

import "io"

func WrapToWriteCloser(writer io.Writer) io.WriteCloser {
	if writeCloser, ok := writer.(io.WriteCloser); ok {
		return writeCloser
	}
	return NopWriteCloser(writer)
}

func WrapToReadCloser(reader io.Reader) io.ReadCloser {
	if readCloser, ok := reader.(io.ReadCloser); ok {
		return readCloser
	}
	return NopReadCloser(reader)
}
