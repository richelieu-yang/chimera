package cookieKit

import (
	"net/http"
	"net/url"
)

// NewCookie
/*
@param path			可以为""
@param domain		可以为""
@param maxAge		单位为秒（s）
@param secure
@param httpOnly
@param sameSite		e.g. http.SameSiteDefaultMode
*/
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
