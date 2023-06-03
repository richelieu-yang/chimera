package errorKit

import "github.com/gogf/gf/v2/errors/gerror"

var (
	// New 用于创建一个自定义错误信息的error对象，并包含堆栈信息。
	New  func(text string) error                        = gerror.New
	Newf func(format string, args ...interface{}) error = gerror.Newf

	// Wrap 用于包裹其他错误error对象，构造成多级的错误信息，包含堆栈信息。
	Wrap  func(err error, text string) error                        = gerror.Wrap
	Wrapf func(err error, format string, args ...interface{}) error = gerror.Wrapf

	// NewSkip 用于创建一个自定义错误信息的error对象，并且忽略部分堆栈信息（按照当前调用方法位置往上忽略）。高级功能，一般开发者很少用得到。
	/*
		@param skip (1) >=0
					(2) 0: 等价于 New || Newf
					(2) 1: 跳过1层（e.g. assert工具类）
	*/
	NewSkip  func(skip int, text string) error                        = gerror.NewSkip
	NewSkipf func(skip int, format string, args ...interface{}) error = gerror.NewSkipf
)
