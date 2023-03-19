package urlKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	"net/url"
)

// EncodeURIComponent 编码
/*
e.g.
("") => ""
*/
func EncodeURIComponent(text string) string {
	return url.QueryEscape(text)
}

// DecodeURIComponent 解码
/*
e.g.
("") => "", nil
*/
func DecodeURIComponent(text string) (string, error) {
	return url.QueryUnescape(text)
}

func ToQueryString(m map[string]string) string {
	var str string

	for k, v := range m {
		if strKit.IsAllNotEmpty(k, v) {
			if strKit.IsNotEmpty(str) {
				str += "&"
			}
			str += k + "=" + EncodeURIComponent(v)
		}
	}
	return str
}

func CombineQueryString(strings ...string) string {
	var qs string

	for _, str := range strings {
		if strKit.IsNotEmpty(str) {
			if strKit.IsNotEmpty(qs) {
				qs += "&"
			}
			qs += str
		}
	}
	return qs
}
