package logrusKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/sirupsen/logrus"
)

// SetConsoleLevel 设置控制台输出的日志级别
func SetConsoleLevel(level logrus.Level) {
	logrus.SetLevel(level)
}

func SetLevel(logger *logrus.Logger, level logrus.Level) {
	if logger != nil {
		logger.SetLevel(level)
	}
}

func StringToLevel(str string) logrus.Level {
	switch strKit.ToLower(str) {
	case "trace":
		return logrus.TraceLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	case "debug":
		fallthrough
	default:
		return logrus.DebugLevel
	}
}
