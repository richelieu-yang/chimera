package assertKit

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

func AssertEmail(email string) error {
	validate := validator.New()
	err := validate.Var(email, "required,email")
	if err != nil {
		return errorKit.SimpleWithExtraSkip(1, "[%s] email(%s) is invalid with error(%s)", funcKit.GetFuncName(1), email, err.Error())
	}
	return nil
}

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

// AssertIPv4
/*
e.g.
	("192.168.9.254")	=> nil
	("127.0.0.1")	 	=> nil
	("localhost")	 	=> test/test1.go:11|main [Assertion failed] ipv4(localhost) is invalid(Key: '' Error:Field validation for '' failed on the 'ipv4' tag)
	("::1")	 			=> test/test1.go:12|main [Assertion failed] ipv4(::1) is invalid(Key: '' Error:Field validation for '' failed on the 'ipv4' tag)
*/
func AssertIPv4(ipv4 string) error {
	validate := validator.New()
	err := validate.Var(ipv4, "required,ipv4")
	if err != nil {
		return errorKit.SimpleWithExtraSkip(1, "[%s] ipv4(%s) is invalid with error(%s)", funcKit.GetFuncName(1), ipv4, err.Error())
	}
	return nil
}
