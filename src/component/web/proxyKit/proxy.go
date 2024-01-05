package proxyKit

import (
	"net/http"
)

// Proxy 代理请求（反向代理，请求转发）.
/*
PS: 转发请求前如果想变更请求头(Request Header)，可以在调用此函数前设置请求头.

@param w		e.g.ctx.Writer
@param r 		e.g.ctx.Request
@param options 	可以不传
*/
func Proxy(w http.ResponseWriter, r *http.Request, targetAddr string, options ...ProxyOption) error {
	opts := loadOptions(options...)
	return opts.proxy(w, r, targetAddr)
}
