package ptrKit

import "github.com/richelieu42/chimera/v2/src/core/errorKit"

// AssertNotNilPointer 断言传参 (1)值非nil; (2)类型为指针
func AssertNotNilPointer(ptr interface{}, paramName string) error {
	if ptr == nil {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] param(name: %s) == nil", paramName)
	}
	if !IsPointer(ptr) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] type(%T) of param(type: %s) isn't pointer", ptr, paramName)
	}
	return nil
}
