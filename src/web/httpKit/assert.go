package httpKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

// AssertHttpUrl 以 "http://" 或 "https://" 开头.
/*
e.g.
	("https://github.com/go-playground/validator") 	=> nil
	("http://github.com/go-playground/validator") 	=> nil
	("ftp://github.com/go-playground/validator") 	=> test/test1.go:11|main [Assertion failed] url(ftp://github.com/go-playground/validator) is invalid(Key: '' Error:Field validation for '' failed on the 'http_url' tag)
*/
func AssertHttpUrl(httpUrl string) error {
	if err := validateKit.HttpUrl(httpUrl); err != nil {
		return errorKit.NewSkip(1, "[%s] httpUrl(%s) is because of with error(%s)", funcKit.GetFuncName(1), httpUrl, err.Error())
	}
	return nil
}
