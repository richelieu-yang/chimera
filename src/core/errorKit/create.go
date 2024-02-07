package errorKit

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/richelieu-yang/chimera/v3/src/funcKit"
)

// New 用于创建一个自定义错误信息的error对象，并包含堆栈信息.
func New(format string, args ...interface{}) error {
	skip := 1
	return gerror.NewSkipf(skip, funcKit.AddEntireCaller(skip+1, format), args...)
}

// NewSkip 用于创建一个自定义错误信息的error对象，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。高级功能，一般开发者很少用得到.
/*
	@param skip (1) >=0
				(2) 0: 等价于 New || New
				(2) 1: 跳过1层（e.g. assert工具类）
*/
func NewSkip(skip int, format string, args ...interface{}) error {
	skip++
	return gerror.NewSkipf(skip, funcKit.AddEntireCaller(skip+1, format), args...)
}

// Wrap 用于包裹其他错误error对象，构造成多级的错误信息，包含堆栈信息.
func Wrap(err error, format string, args ...interface{}) error {
	skip := 1
	return gerror.WrapSkipf(skip, err, funcKit.AddEntireCaller(skip+1, format), args...)
}
