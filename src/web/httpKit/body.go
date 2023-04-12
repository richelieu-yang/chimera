package httpKit

import (
	"bytes"
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"io"
	"net/http"
)

// MakeRequestBodySeekable
/*
PS:
(1) 某个路由涉及代理（请求转发）的话，需要在handler里面首先调用此方法.
(2) 与 Proxy() 搭配使用.
*/
func MakeRequestBodySeekable(r *http.Request) error {
	if r.Body == nil {
		return nil
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	readSeeker := bytes.NewReader(data)
	r.Body = ioKit.NopReadSeekCloser(readSeeker)
	return nil
}
