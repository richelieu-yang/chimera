package httpKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"net/http"
)

// GetAcceptLanguages 根据请求头"Accept-Language".
func GetAcceptLanguages(req *http.Request) []string {
	langs := make([]string, 0)

	headerValue := GetHeader(req.Header, "Accept-Language")
	str := strKit.Split(headerValue, ";")[0]

	str = strKit.TrimSpace(str)
	s := strKit.Split(str, ",")
	for _, lang := range s {
		if strKit.IsEmpty(lang) {
			continue
		}
		langs = append(langs, lang)
	}
	return langs
}
