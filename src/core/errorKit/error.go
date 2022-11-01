// Package errorKit
/*
极客时间的"Go error处理最佳实践": https://u.geekbang.org/lesson/376?article=513711

输出带有堆栈的error时，建议使用: logrusKit.PrintError 或 logrusKit.PrintErrorWithLogger.

使用时的注意点:
(a) 在应用代码（业务代码）中，使用 errorKit.New 返回错误；
(b) 在项目工程里面，如果调用其它包内的函数，通常简单的直接返回error，即向上抛;
(c) 如果和其它库（第三方库、标准库）进行协作，考虑使用 errorKit.Wrap 保存堆栈信息；
(d) 不要到处打日志（比如每个错误产生的地方），直接往上抛的情况不需要打日志；
(e) 在 程序的顶部 或 工作的goroutine顶部（请求入口），使用 %+v 输出堆栈详情.
(f) 使用 errorKit.Cause 获取root error，再和sentinel error判定.
(g) 如果你的函数不打算处理这个错误，你没法处理，那么应该携带上足够多的上下文，然后往上抛；（Wrap errors）
(h) 如果错误被处理了，那么它就不应该被往上抛并且记日志，

如果是 第三方库 或 基础库（被很多人使用并且是跨很多项目的），
(a) 应该返回 root error，它的原始错误是什么就是什么；
(b) 如果返回的新的error，应该使用 errorKit.Simple 来创建error；
(c) errorKit.New、errorKit.Wrap 和 errorKit.WithMessage: 业务代码可以使用，第三方库和基础库不应该使用（以避免保存多次堆栈信息，太冗余了）；
*/
package errorKit

import (
	"errors"
	"fmt"
	errors2 "github.com/pkg/errors"
	"github.com/richelieu42/go-scales/src/funcKit"
)

// Simple 新建error（指针；不会携带堆栈信息）
func Simple(format string, args ...interface{}) error {
	format = funcKit.AddFuncInfoToString(format, 1)
	return errors.New(fmt.Sprintf(format, args...))
}

// SimpleWithExtraSkip
/*
@param extraSkip 额外跳过的层数（>= 0）
*/
func SimpleWithExtraSkip(extraSkip int, format string, args ...interface{}) error {
	format = funcKit.AddFuncInfoToString(format, 1)
	return errors.New(fmt.Sprintf(format, args...))
}

// New 新建error（指针；会携带堆栈信息）
/*
@param format ！！！推荐使用格式"{包名}: {错误的具体信息}"，e.g. "scales-strKit: param is invalid"
@return 每次都会返回一个新的指针

e.g. 返回值为什么是指针（详见 errors.Errorf）
err := errorKit.Simple("123")
err1 := errorKit.Simple("123")
fmt.Println(err == err)  	// true
fmt.Println(err == err1) 	// false（因为内存地址不一样）
*/
func New(format string, args ...interface{}) error {
	format = funcKit.AddFuncInfoToString(format, 1)
	return errors2.Errorf(format, args...)
}

//// NewWithExtraSkip
///*
//@param extraSkip 额外跳过的层数
//*/
//func NewWithExtraSkip(extraSkip int, format string, args ...interface{}) error {
//	format = addCallerInfoToFormat(format, extraSkip)
//	return errors2.Errorf(format, args...)
//}

// IsErrorWithStack 错误是否有堆栈信息？
/*
不考虑复杂情况（e.g. 混用 errorKit.Wrap 和 errorKit.WithMessage），
	errorKit.Wrap 和 errorKit.New 返回的error（非nil），将返回: true；
	errorKit.WithMessage 返回的error（非nil），将返回: false.

e.g.
err := io.EOF
err1 := errorKit.Wrap(io.EOF, "1")
err2 := errorKit.WithMessage(io.EOF, "2")
err3 := errorKit.New("3")

fmt.Println(errorKit.IsErrorWithStack(err))  // false
fmt.Println(errorKit.IsErrorWithStack(err1)) // true
fmt.Println(errorKit.IsErrorWithStack(err2)) // false
fmt.Println(errorKit.IsErrorWithStack(err3)) // true
*/
func IsErrorWithStack(err error) bool {
	type stackTracer interface {
		StackTrace() errors2.StackTrace
	}

	if err == nil {
		return false
	}
	_, ok := err.(stackTracer)
	return ok
}

func WithStack(err error) error {
	return errors2.WithStack(err)
}

// Wrap
/*
Wrap returns an error annotating err with a stack trace at the point Wrap is called, and the supplied message.
If err is nil, Wrap returns nil.

@param format 建议首字母小写，且最后面不要加标点符号，否则"%v"输出会比较难看（%v: message在左边；%+v: message在下边）
*/
func Wrap(err error, format string, args ...interface{}) error {
	format = funcKit.AddFuncInfoToString(format, 1)
	return errors2.Wrapf(err, format, args...)
}

// WithMessage
/*
WithMessage annotates err with a new message.
If err is nil, WithMessage returns nil.
%v：	message在左边；
%+v：	message在下面。
*/
func WithMessage(err error, format string, args ...interface{}) error {
	format = funcKit.AddFuncInfoToString(format, 1)
	return errors2.WithMessagef(err, format, args...)
}

// Cause 获取传参err的root error（错误的根因）.
/*
(1) If the error does not implement Cause, the original error will be returned.（返回最原始的错误）
(2) If the error is nil, nil will be returned without further investigation.

e.g.
(nil) => nil
*/
func Cause(err error) error {
	return errors2.Cause(err)
}
