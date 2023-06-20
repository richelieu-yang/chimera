package errorKit

import (
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	// Is reports whether any error in err's tree matches target.
	Is func(err, target error) bool = errors.Is

	// Equal 错误比较
	Equal func(err, target error) bool = gerror.Equal

	// HasStack 判断错误是否带堆栈
	HasStack func(err error) bool = gerror.HasStack

	// Stack 获取堆栈信息
	Stack func(err error) string = gerror.Stack

	// Current 获取当前error
	Current func(err error) error = gerror.Current

	// Unwrap 获取层级错误的下一级错误error接口对象(当下一层级不存在时，返回nil)
	Unwrap func(err error) error = gerror.Unwrap

	// Cause 获取根错误error
	Cause func(err error) error = gerror.Cause
)
