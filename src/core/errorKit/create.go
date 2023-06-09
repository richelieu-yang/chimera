package errorKit

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

// New 用于创建一个自定义错误信息的error对象，并包含堆栈信息.
func New(text string) error {
	return gerror.NewSkip(1, addFuncInfo(text))
}

func Newf(format string, args ...interface{}) error {
	return gerror.NewSkipf(1, addFuncInfo(format), args...)
}

// NewSkip 用于创建一个自定义错误信息的error对象，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。高级功能，一般开发者很少用得到.
/*
	@param skip (1) >=0
				(2) 0: 等价于 New || Newf
				(2) 1: 跳过1层（e.g. assert工具类）
*/
func NewSkip(skip int, text string) error {
	return gerror.NewSkip(skip+1, addFuncInfo(text))
}

func NewSkipf(skip int, format string, args ...interface{}) error {
	return gerror.NewSkipf(skip+1, addFuncInfo(format), args...)
}

// Wrap 用于包裹其他错误error对象，构造成多级的错误信息，包含堆栈信息.
func Wrap(err error, text string) error {
	return gerror.WrapSkip(1, err, addFuncInfo(text))
}

func Wrapf(err error, format string, args ...interface{}) error {
	return gerror.WrapSkipf(1, err, addFuncInfo(format), args...)
}

func addFuncInfo(text string) string {
	return funcKit.GetEntireCaller(3) + " " + text
}
