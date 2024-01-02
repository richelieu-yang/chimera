package httpKit

import "net/http"

// GetReferer
/*
PS:
(1) 涉及 Referer验证;
(2) 参考: notes/Web（漏洞等）/Web.wps
*/
func GetReferer(r *http.Request) string {
	return r.Referer()
}
