package proxyKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Proxy 代理请求（反向代理，请求转发）.
/*
PS: 转发请求前如果想变更请求头(Request Header)，可以在调用此函数前设置请求头.

@param w			e.g.ctx.Writer
@param r 			e.g.ctx.Request
@param targetAddr	e.g."127.0.0.1:12345"
@param options 		可以不传
*/
func Proxy(w http.ResponseWriter, r *http.Request, targetAddr string, options ...ProxyOption) error {
	opts := loadOptions(options...)

	return opts.proxy(w, r, targetAddr)
}

func ProxyWithGin(ctx *gin.Context, targetAddr string, options ...ProxyOption) error {
	opts := loadOptions(options...)
	opts.ctx = ctx

	return opts.proxy(ctx.Writer, ctx.Request, targetAddr)
}
