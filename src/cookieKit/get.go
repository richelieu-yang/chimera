package cookieKit

import (
	"net/http"
	"net/url"
)

// GetCookie
/*
!!!:
(1) 获取到的 *http.Cookie 实例，如果对其进行了修改，浏览器端并不会同步修改（除非修改后set回去）.
(2) 如果不存在与传参name对应的cookie，将返回 (nil, http.ErrNoCookie).

@param req 	e.g. gin.Context Request
@param name	可以为""，但无意义
*/
func GetCookie(req *http.Request, name string) (*http.Cookie, error) {
	return req.Cookie(name)
}

// GetCookieValue
/*
@param req e.g. gin.Context Request
@param name	可以为""，但无意义
*/
func GetCookieValue(req *http.Request, name string) (string, error) {
	cookie, err := GetCookie(req, name)
	if err != nil {
		return "", err
	}

	// 解码
	value, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return "", err
	}
	return value, nil
}
