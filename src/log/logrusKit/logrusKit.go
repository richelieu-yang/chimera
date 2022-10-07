package logrusKit

import (
	"gitee.com/richelieu042/go-scales/src/core/timeKit"
	"gitee.com/richelieu042/go-scales/src/mainControl"
	"github.com/sirupsen/logrus"
)

// Initialize
/*
PS: logrus依赖的默认日志级别为 INFO.
*/
func Initialize(level logrus.Level, timestampFormat timeKit.TimeFormat) {
	// 日志级别
	SetConsoleLevel(level)

	// 时间戳格式
	formatter := NewTextFormatter(timestampFormat)
	logrus.SetFormatter(formatter)
}

// InitializeByDefault 默认地初始化 logrus
/*
！！！：
(1) 不要单独使用"\r"，否则输出会有问题；
(2) "\r\n"一起用没问题.

e.g. 反例
logrusKit.InitializeByDefault()
logrus.Info("--------------------------------")
logrus.Infof("os:\r%s", osKit.OS)
logrus.Info("--------------------------------")
*/
func InitializeByDefault() {
	Initialize(logrus.DebugLevel, timeKit.CommonFormat)
}

func OnDebugChanged() {
	logrus.SetLevel(mainControl.GetLogrusLevel())
}
