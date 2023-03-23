package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

func main() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (funcName string, fileName string) {
			s := strings.Split(f.Function, ".")
			funcName = s[len(s)-1]

			s1 := strKit.Split(f.File, "/")
			length := len(s1)
			if length >= 2 {
				fileName = fmt.Sprintf("%s/%s:%d", s1[length-2], s1[length-1], f.Line)
			} else {
				fileName = fmt.Sprintf("%s:%d", f.File, f.Line)
			}

			return funcName, fileName
		},
	}
	l.Info("测试")
}
