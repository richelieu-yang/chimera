package httpKit

import (
	"bytes"
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"io"
	"net/http"
)

// MakeRequestBodySeekable
/*
与 Proxy() 搭配使用.
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
