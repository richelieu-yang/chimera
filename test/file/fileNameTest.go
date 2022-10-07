package main

import (
	"gitee.com/richelieu042/go-scales/src/core/file/fileKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info(fileKit.GetName("c:/d/test.DOCX"))

	logrus.Infof("%s", fileKit.GetPrefix("c:/d/test"))
	logrus.Infof("%s", fileKit.GetSuffix("c:/d/test"))

	logrus.Infof("%s", fileKit.GetPrefix("c:/d/test.DOCX"))
	logrus.Infof("%s", fileKit.GetSuffix("c:/d/test.DOCX"))
}
