package assertKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/ptrKit"
)

// Pointer 断言是指针
/*
@param ptr			可以为nil，但会断言失败
@param variableName 传参ptr在上一层的变量名
*/
func Pointer(ptr interface{}, variableName string) error {
	if !ptrKit.IsPointer(ptr) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable(name: %s, type: %T) isn't a pointer", variableName, ptr)
	}
	return nil
}
