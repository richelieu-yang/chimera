package ptrKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

// AssertNotNilPointer 断言传参 (1)值非nil; (2)类型为指针
func AssertNotNilPointer(ptr interface{}) error {
	if ptr == nil {
		return errorKit.SimpleWithExtraSkip(1, "[%s] ptr == nil", funcKit.GetFuncName(1))
	}
	if !IsPointer(ptr) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] ptr(type: %T) isn't a pointer", funcKit.GetFuncName(1), ptr)
	}
	return nil
}
