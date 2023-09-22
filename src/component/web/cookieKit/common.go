package cookieKit

import (
	"net/http"
	"time"
)

// SetCookieMaxAge 修改Cookie的MaxAge属性（兼容IE浏览器）
/*
PS:
(1) 同时设置 MaxAge 和 Expires 的情况下，
	(a) IE浏览器，忽略 MaxAge，使用 Expires;
	(b) 其它浏览器，忽略 Expires，使用 MaxAge.
(2) IE浏览器只认 Expires，但这样有缺陷: 浏览器（客户端）和服务端存在时间差的情况下，Cookie可能会提前或延后过期.
*/
func SetCookieMaxAge(cookie *http.Cookie, maxAge int) {
	cookie.MaxAge = maxAge

	/* 兼容IE浏览器（它不支持MaxAge） */
	if maxAge > 0 {
		d := time.Duration(maxAge) * time.Second
		cookie.Expires = time.Now().Add(d)
	} else if maxAge < 0 {
		// Set it to the past to expire now.（即删除cookie）
		cookie.Expires = time.Unix(1, 0)
	}
}
