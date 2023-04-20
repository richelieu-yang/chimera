package urlKit

import "net/url"

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
