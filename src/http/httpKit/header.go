// Package httpKit
/*
请求头 || 响应头
key: 不区分大小写
*/
package httpKit

import (
	"net/http"
)

// AddHeader
/*
e.g.
	header := make(map[string][]string)

	fmt.Println(header)					// map[]
	httpKit.AddHeader(header, "k", "0")
	fmt.Println(header)					// map[K:[0]]
	httpKit.AddHeader(header, "k", "1")
	fmt.Println(header)					// map[K:[0 1]]
*/
func AddHeader(header http.Header, key, value string) {
	header.Add(key, value)
}

// SetHeader
/*
e.g.
	header := make(map[string][]string)

	fmt.Println(header) 				// map[]
	httpKit.AddHeader(header, "k", "0")
	fmt.Println(header) 				// map[K:[0]]
	httpKit.AddHeader(header, "k", "1")
	fmt.Println(header) 				// map[K:[0 1]]

	httpKit.SetHeader(header, "k", "2")
	fmt.Println(header) 				// map[K:[2]]
	httpKit.SetHeader(header, "k", "")
	fmt.Println(header) 				// map[]
*/
func SetHeader(header http.Header, key, value string) {
	if value == "" {
		DelHeader(header, key)
		return
	}
	header.Set(key, value)
}

func GetHeader(header http.Header, key string) string {
	return header.Get(key)
}

func DelHeader(header http.Header, key string) {
	header.Del(key)
}

// SetCacheControlNoCache 实际上是有缓存的）浏览器对请求回来的response做缓存，但是每次在向客户端（浏览器）提供响应数据时，缓存都要向服务器评估缓存响应的有效性。
/*
PS: 详见"Web.docx".
*/
func SetCacheControlNoCache(header http.Header) {
	SetHeader(header, "Cache-Control", "no-cache")
}

// SetCacheControlNoStore 禁止一切缓存.
/*
PS: 详见"Web.docx".
*/
func SetCacheControlNoStore(header http.Header) {
	SetHeader(header, "Cache-Control", "no-store")
}

// GetUserAgent 获取http请求头中"User Agent"的值.
/*
参考: https://www.sunzhongwei.com/golang-gin-for-user-agent-in-http-request-header-value

e.g.
Chrome浏览器: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36
Safari浏览器: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.5 Safari/605.1.15
*/
func GetUserAgent(header http.Header) string {
	return GetHeader(header, "User-Agent")
}

// GetOrigin 获取请求的Origin
func GetOrigin(header http.Header) string {
	return GetHeader(header, "Origin")
}
