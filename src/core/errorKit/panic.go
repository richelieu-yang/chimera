package errorKit

import (
	"github.com/richelieu42/go-scales/src/funcKit"
	"github.com/sirupsen/logrus"
)

// PanicByError
/*
PS: 如果传参只有1个error实例，可以使用此方法.
*/
func PanicByError(err error) {
	logrus.Panic(err)
}

// Panic
/*

e.g. 想要输出error的情况
传参: ("fail to initialize %s, error:\n%+v", "gin", err)
*/
func Panic(format string, args ...interface{}) {
	format = funcKit.AddFuncInfoToString(format, 1)
	logrus.Panicf(format, args...)
}
