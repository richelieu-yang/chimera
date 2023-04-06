package httpKit

import (
	"io"
	"sync/atomic"
)

type Repeat struct {
	Reader io.ReaderAt
	Offset int64
}

// Read 重写读方法，使每次读取request.Body时能从指定位置读取
func (p *Repeat) Read(val []byte) (n int, err error) {
	n, err = p.Reader.ReadAt(val, p.Offset)
	atomic.AddInt64(&p.Offset, int64(n))
	return
}

// Reset 重置偏移量
func (p *Repeat) Reset() {
	atomic.StoreInt64(&p.Offset, 0)
}

func (p *Repeat) Close() error {
	// 因为req.Body实现了readcloser接口，所以要实现close方法
	// 但是repeat中的值有可能是只读的，所以这里只是尝试关闭一下。
	if p.Reader != nil {
		if rc, ok := p.Reader.(io.Closer); ok {
			return rc.Close()
		}
	}
	return nil
}
