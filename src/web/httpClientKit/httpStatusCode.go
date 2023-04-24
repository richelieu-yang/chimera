package httpClientKit

import (
	"net/http"
)

// IsCodeValid 响应的http状态码 是否有效？
/*
PS: 判断参考了jQuery.
*/
func IsCodeValid(code int) bool {
	return code >= 200 && code < 300 || code == http.StatusNotModified
}
