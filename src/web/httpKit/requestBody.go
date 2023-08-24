package httpKit

import (
	"bytes"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"io"
	"net/http"
)

// MakeRequestBodySeekable
/*
Deprecated: Gin不需要调用此方法.

PS:
(1) 一般与 proxy() 搭配使用.
(2) 某个路由涉及代理（请求转发）的话，需要在handler里面首先调用此方法.
*/
func MakeRequestBodySeekable(req *http.Request) error {
	if req.Body == nil {
		return nil
	}

	if _, ok := req.Body.(io.Seeker); ok {
		// 已经实现了 io.Seeker，避免重复调用
		return nil
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	readSeeker := bytes.NewReader(data)
	req.Body = ioKit.NopCloserToReadSeeker(readSeeker)
	return nil
}

// ResetRequestBody 重置请求体，以防: 已经读完body了，请求转发给别人，别人收到的请求没内容.
/*
PS: req.Body可以为nil.
*/
func ResetRequestBody(req *http.Request) (bool, error) {
	if req.Body == nil {
		return true, nil
	}

	seeker, ok := req.Body.(io.Seeker)
	if ok {
		_, err := seeker.Seek(0, io.SeekStart)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
