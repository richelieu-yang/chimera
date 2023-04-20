package urlKit

import (
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

// ToQueryString
/*
@param m 会对值进行 编码 操作
@return 可能为""

e.g.
	(nil) => ""
*/
func ToQueryString(queryParams map[string]string) string {
	var str string

	for k, v := range queryParams {
		// PS: k和v都有可能是""
		if strKit.IsNotEmpty(str) {
			str += "&"
		}
		if strKit.IsNotEmpty(v) {
			str += k + "=" + EncodeURIComponent(v)
		} else {
			str += k
		}
	}
	return str
}

// AttachQueryParamsToUrl
/*
@param url 也可以是: Request.URL.RawQuery
*/
func AttachQueryParamsToUrl(url string, queryParams map[string]string) string {
	queryStr := ToQueryString(queryParams)
	if strKit.IsEmpty(queryStr) {
		return url
	}

	i := strKit.Index(url, "?")
	if i == -1 {
		return url + "?" + queryStr
	} else if i == len(url)-1 {
		// 传参url最后一个字符为"?"
		return url + queryStr
	}
	return url + "&" + queryStr
}
