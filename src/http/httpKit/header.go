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
