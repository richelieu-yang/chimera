package ioKit

import (
	"io"
	"sync"
)

type multiWriterCloser struct {
	lock         *sync.Mutex
	writeClosers []io.WriteCloser
}

func (multi *multiWriterCloser) Write(p []byte) (n int, err error) {
	multi.lock.Lock()
	defer multi.lock.Unlock()

	for _, w := range multi.writeClosers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func (multi *multiWriterCloser) WriteString(s string) (n int, err error) {
	multi.lock.Lock()
	defer multi.lock.Unlock()

	var p []byte // lazily initialized if/when needed
	for _, w := range multi.writeClosers {
		if sw, ok := w.(io.StringWriter); ok {
			n, err = sw.WriteString(s)
		} else {
			if p == nil {
				p = []byte(s)
			}
			n, err = w.Write(p)
		}
		if err != nil {
			return
		}
		if n != len(s) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(s), nil
}

func (multi *multiWriterCloser) Close() (err error) {
	multi.lock.Lock()
	defer multi.lock.Unlock()

	defer func() {
		multi.writeClosers = nil
	}()

	return CloseWriteClosers(multi.writeClosers...)
}

// MultiWriteCloser
/*
PS: 参考了 io.MultiWriter().
*/
func MultiWriteCloser(writeClosers ...io.WriteCloser) io.WriteCloser {
	all := make([]io.WriteCloser, 0, len(writeClosers))

	for _, writeCloser := range writeClosers {
		if multi, ok := writeCloser.(*multiWriterCloser); ok {
			all = append(all, multi.writeClosers...)
		} else {
			all = append(all, writeCloser)
		}
	}
	return &multiWriterCloser{
		lock:         new(sync.Mutex),
		writeClosers: all,
	}
}
