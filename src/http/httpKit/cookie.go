package httpKit

import "net/http"

// ClearCookies
// 清空Cookie.
func ClearCookies(req *http.Request) {
	header := req.Header
	if header != nil {
		delete(header, "Cookie")
	}
}
