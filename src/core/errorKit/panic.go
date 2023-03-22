package errorKit

import (
	"github.com/richelieu42/chimera/src/funcKit"
	"github.com/sirupsen/logrus"
)

func PanicByError(err error) {
	logrus.Panic(err)
}

func Panic(format string, args ...interface{}) {
	format = funcKit.AddFuncInfoToString(format, 1)
	logrus.Panicf(format, args...)
}
