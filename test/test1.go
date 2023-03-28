package main

import (
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/logrusx"
)

func main() {
	logx.Info("---")

	writer := logrusx.NewLogrusWriter(func(logger *logrus.Logger) {
		formatter := logrusKit.NewTextFormatter("")
		logger.SetFormatter(formatter)
		//logger.SetFormatter(&logrus.TextFormatter{})
	})
	logx.SetWriter(writer)

	logx.Info("+++")
}
