package gozeroKit

import (
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/zero-contrib/logx/logrusx"
)

// SetLogrusWriter go-zero与第三方日志库logrus集成
/*
@param formatter 可以为nil
*/
func SetLogrusWriter(formatter logrus.Formatter) {
	if formatter == nil {
		formatter = logrusKit.DefaultTextFormatter
	}

	writer := logrusx.NewLogrusWriter(func(logger *logrus.Logger) {
		logger.SetFormatter(formatter)
	})
	logx.SetWriter(writer)
}
