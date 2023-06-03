package httpKit

import (
	"bytes"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"io"
	"net/http"
)

// MakeRequestBodySeekable
/*
PS:
(1) 一般与 proxy() 搭配使用.
(2) 某个路由涉及代理（请求转发）的话，需要在handler里面首先调用此方法.

*/
func MakeRequestBodySeekable(r *http.Request) error {
	if r.Body == nil {
		return nil
	}
	if _, ok := r.Body.(io.Seeker); ok {
		// 已经实现了 io.Seeker，避免重复调用
		return nil
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	readSeeker := bytes.NewReader(data)
	r.Body = ioKit.NopCloserToReadSeeker(readSeeker)
	return nil
}
