package mainControl

import (
	"github.com/sirupsen/logrus"
)

var debug = false

func SetDebug(flag bool) {
	debug = flag
}

func IsDebug() bool {
	return debug
}

// GetLogrusLevel 获取logrus的日志级别（默认：INFO）
func GetLogrusLevel() logrus.Level {
	if IsDebug() {
		return logrus.DebugLevel
	}
	return logrus.InfoLevel
}
