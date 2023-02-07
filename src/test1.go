package main

import (
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var (
	Author    string
	GoVersion string
)

func main() {
	logrusKit.InitializeByDefault()

	logrus.Infof("Author: [%s].", Author)
	logrus.Infof("GoVersion: [%s].", GoVersion)
}
