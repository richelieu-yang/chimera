package httpKit

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

// AssertHttpUrl 以 "http://" 或 "https://" 开头.
/*
e.g.
	("https://github.com/go-playground/validator") 	=> nil
	("http://github.com/go-playground/validator") 	=> nil
	("ftp://github.com/go-playground/validator") 	=> test/test1.go:11|main [Assertion failed] url(ftp://github.com/go-playground/validator) is invalid(Key: '' Error:Field validation for '' failed on the 'http_url' tag)
*/
func AssertHttpUrl(httpUrl string) error {
	validate := validator.New()
	err := validate.Var(httpUrl, "required,http_url")
	if err != nil {
		return errorKit.SimpleWithExtraSkip(1, "[%s] httpUrl(%s) is invalid with error(%s)", funcKit.GetFuncName(1), httpUrl, err.Error())
	}
	return nil
}
