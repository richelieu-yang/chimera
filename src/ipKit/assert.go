package ipKit

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

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
