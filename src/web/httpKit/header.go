// Package httpKit
/*
请求头 || 响应头
key: 不区分大小写
*/
package httpKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"net/http"
)

// HeaderToMap http.Header（即 map[string][]string） => map[string]interface{}
func HeaderToMap(header http.Header) map[string]interface{} {
	return mapKit.MapValues(header, func(value []string, key string) interface{} {
		return value
	})
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

// GetHeader
/*
PS: 存在多个值的话（[]string），返回第一个值.

@param key 不存在对应值的话，将返回 ""
*/
func GetHeader(header http.Header, key string) string {
	return header.Get(key)
}

// GetHeaderValues
/*
@param key 不存在对应值的话，将返回 nil
*/
func GetHeaderValues(header http.Header, key string) []string {
	return header.Values(key)
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

// AddHeader
/*
e.g.
	header := make(map[string][]string)

	AddHeader(header, "k", "0")
	fmt.Println(header) 		// map[K:[0]]
	AddHeader(header, "k", "1")
	fmt.Println(header) 		// map[K:[0 1]]
	AddHeader(header, "k", "1")
	fmt.Println(header) 		// map[K:[0 1 1]]
*/
func AddHeader(header http.Header, key, value string) {
	header.Add(key, value)
}

// AddHeaderIfMissing
/*
PS: 区分大小写.
*/
func AddHeaderIfMissing(header http.Header, key, value string) {
	if !HeaderContainsValue(header, key, value) {
		header.Add(key, value)
	}
}

// AddHeaderIfMissingIgnoreCase
/*
PS: 不区分大小写.
*/
func AddHeaderIfMissingIgnoreCase(header http.Header, key, value string) {
	if !HeaderContainsValueIgnoreCase(header, key, value) {
		header.Add(key, value)
	}
}

// HeaderContainsValue Header中，指定key对应的value切片是否包含指定value？
/*
PS: 区分大小写.
*/
func HeaderContainsValue(header http.Header, key, value string) bool {
	values := header.Values(key)
	return sliceKit.Contains(values, value)
}

// HeaderContainsValueIgnoreCase Header中，指定key对应的value切片是否包含指定value？\
/*
PS: 不区分大小写.
*/
func HeaderContainsValueIgnoreCase(header http.Header, key, value string) bool {
	values := header.Values(key)
	return sliceKit.ContainsBy(values, func(item string) bool {
		return strKit.EqualsIgnoreCase(item, value)
	})
}
