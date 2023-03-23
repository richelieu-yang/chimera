package main

import (
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
)

func main() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			_, fileName := path.Split(f.File)
			return funcName, fileName
		},
	}
	l.Info("example of custom format caller")
}
