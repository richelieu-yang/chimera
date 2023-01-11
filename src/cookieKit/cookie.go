package cookieKit

import (
	"net/http"
	"net/url"
)

func NewCookie(name, value, path, domain string, maxAge int, secure, httpOnly bool, sameSite http.SameSite) *http.Cookie {
	// 参考: gin.Context.SetCookie()
	if path == "" {
		path = "/"
	}

	return &http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		Path:     path,
		Domain:   domain,
		MaxAge:   maxAge,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: sameSite,
	}
}

// GetCookie
/*
@param req e.g. gin.Context Request
*/
func GetCookie(req *http.Request, name string) (*http.Cookie, error) {
	return req.Cookie(name)
}

// GetCookieValue
/*
@param req e.g. gin.Context Request
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
