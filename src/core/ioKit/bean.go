package ioKit

import "io"

// readSeekCloser 实现了 io.ReadSeekCloser 接口
type readSeekCloser struct {
	readSeeker io.ReadSeeker
}

func (obj *readSeekCloser) Read(p []byte) (n int, err error) {
	return obj.readSeeker.Read(p)
}

func (obj *readSeekCloser) Seek(offset int64, whence int) (int64, error) {
	return obj.readSeeker.Seek(offset, whence)
}

func (obj *readSeekCloser) Close() error {
	// do nothing
	return nil
}
