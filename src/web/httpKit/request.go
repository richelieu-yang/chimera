package httpKit

import (
	"net/http"
)

// GetRoute 获取: 路由（不带query）.
/*
e.g.
	http://127.0.0.1/a/b?1=1&2=2 => "/a/b"
*/
func GetRoute(req *http.Request) string {
	return req.URL.Path
}

// GetRouteWithQuery 获取: 路由（带query）.
/*
e.g.
	http://127.0.0.1/a/b?1=1&2=2 => "/a/b?1=1&2=2"
*/
func GetRouteWithQuery(req *http.Request) string {
	return req.RequestURI
}

// GetURLRawQuery
/*
e.g.
http://127.0.0.1/a/b?1=1&2=2 => "1=1&2=2"
*/
func GetURLRawQuery(req *http.Request) string {
	return req.URL.RawQuery
}

// GetProto
/*
@return "HTTP/1.0" || "HTTP/1.1" || ...
*/
func GetProto(req *http.Request) string {
	return req.Proto
}

// GetRequestUrl 返回当前接口的url.
/*
PS: 包括query、fragment.
*/
func GetRequestUrl(req *http.Request) string {
	return req.URL.String()

	//url := req.URL
	//
	///* scheme */
	//var scheme string
	//if strKit.IsEmpty(url.Scheme) {
	//	if websocket.IsWebSocketUpgrade(req) {
	//		scheme = operationKit.Ternary(req.TLS != nil, "wss", "ws")
	//	} else {
	//		scheme = operationKit.Ternary(req.TLS != nil, "https", "http")
	//	}
	//} else {
	//	scheme = url.Scheme
	//}
	//
	///* host */
	//host := url.Host
	//if strKit.IsEmpty(host) {
	//	host = req.Host
	//}
	//
	///* path */
	//path := url.Path
	//
	//return fmt.Sprintf("%s://%s%s", scheme, host, path)
}
