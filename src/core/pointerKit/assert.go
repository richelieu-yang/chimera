package pointerKit

import "github.com/richelieu42/go-scales/src/core/errorKit"

// AssertPointer 断言是指针
/*
@param ptr			可以为nil，但会断言失败
@param variableName 传参ptr在上一层的变量名
*/
func AssertPointer(ptr interface{}, variableName string) error {
	if !IsPointer(ptr) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable %s(type: %T) isn't a pointer", variableName, ptr)
	}
	return nil
}
