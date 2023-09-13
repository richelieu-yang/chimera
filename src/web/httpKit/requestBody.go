package httpKit

import (
	"bytes"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
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
func MakeRequestBodySeekable(req *http.Request) error {
	// 特殊情况: req.Body == http.NoBody，http客户端发的是post请求，但是没有request body（即没post参数）
	if req.Body == nil || req.Body == http.NoBody {
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
	// bytes.NewReader() 的返回值实现了 io.Reader、io.Seeker 接口
	reader := bytes.NewReader(data)
	req.Body = ioKit.NopCloserToReadSeeker(reader)
	return nil
}

// ResetRequestBody 重置请求体，以防: 已经读完body了，请求转发给别人，别人收到的请求没内容.
/*
PS: req.Body可以为nil.
*/
func ResetRequestBody(req *http.Request) error {
	if req.Body == nil || req.Body == http.NoBody {
		// 无需重置
		return nil
	}

	seeker, ok := req.Body.(io.Seeker)
	if !ok {
		// 不能重置
		return errorKit.New("body(%T) is unable to seek", req.Body)
	}
	_, err := seeker.Seek(0, io.SeekStart)
	if err != nil {
		// 重置失败
		return err
	}
	return nil
}
