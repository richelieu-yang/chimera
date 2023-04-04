package cookieKit

import "net/http"

func SetCookieByAttributes(writer http.ResponseWriter, name, value, path, domain string, maxAge int, secure, httpOnly bool, sameSite http.SameSite) {
	cookie := NewCookie(name, value, path, domain, maxAge, secure, httpOnly, sameSite)
	SetCookie(writer, cookie)
}

// SetCookie
/*
@param writer e.g. gin.Context Writer
@param cookie 可以为nil，但这样的话调用此方法就没意义了
*/
func SetCookie(writer http.ResponseWriter, cookie *http.Cookie) {
	if cookie != nil {
		// 参考: gin.Context.SetCookie()
		if cookie.Path == "" {
			cookie.Path = "/"
		}
	}
	http.SetCookie(writer, cookie)
}
