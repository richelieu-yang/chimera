package ioKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"io"
	"os"
	"sync"
)

type writeCloser struct {
	lock   *sync.Mutex
	writer io.Writer
}

func (wc *writeCloser) Write(p []byte) (int, error) {
	if wc == nil {
		return 0, errorKit.Simple("wc == nil")
	}

	wc.lock.Lock()
	defer wc.lock.Unlock()

	return wc.writer.Write(p)
}

func (wc *writeCloser) WriteString(s string) (int, error) {
	if wc == nil {
		return 0, errorKit.Simple("wc == nil")
	}

	wc.lock.Lock()
	defer wc.lock.Unlock()

	return io.WriteString(wc.writer, s)
}

func (wc *writeCloser) Close() error {
	if wc == nil {
		return errorKit.Simple("wc == nil")
	}

	wc.lock.Lock()
	defer wc.lock.Unlock()

	return CloseWriter(wc.writer)
}

// WrapToWriteCloser io.Writer => io.WriteCloser
func WrapToWriteCloser(writer io.Writer) (io.WriteCloser, error) {
	switch writer {
	case nil:
		return nil, errorKit.Simple("writer == nil")
	case os.Stdout:
		fallthrough
	case os.Stderr:
		// 这2种情况必须继续执行以封装，以免被误操作而关闭
	default:
		// 如果本来就是 io.WriteCloser类型，直接返回
		if tmp, ok := writer.(io.WriteCloser); ok {
			return tmp, nil
		}
	}

	return &writeCloser{
		lock:   new(sync.Mutex),
		writer: writer,
	}, nil
}
