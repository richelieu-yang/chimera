package cookieKit

import "net/http"

func DeleteCookieByName(req *http.Request, writer http.ResponseWriter, name string) {
	cookie, err := GetCookie(req, name)
	if err != nil {
		// 不存在对应的cookie，无需删除
		return
	}

	DeleteCookie(writer, cookie)
}

// DeleteCookie
/*
@param cookie 可以为nil
*/
func DeleteCookie(writer http.ResponseWriter, cookie *http.Cookie) {
	if cookie == nil {
		return
	}
	cookie.MaxAge = -1
	SetCookie(writer, cookie)
}
