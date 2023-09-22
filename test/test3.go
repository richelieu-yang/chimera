package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/sirupsen/logrus"
)

type PrefixHook struct {
	Prefix string
}

func (hook *PrefixHook) Fire(entry *logrus.Entry) error {
	if strKit.IsNotEmpty(hook.Prefix) {
		entry.Message = fmt.Sprintf("%s %s", hook.Prefix, entry.Message)
	}
	return nil
}

func (hook *PrefixHook) Levels() []logrus.Level {
	//return logrus.AllLevels

	// 只有 INFO、WARN 级别
	return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}
}

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("Trace")
	logrus.Debug("Debug")
	logrus.Info("Info")
	logrus.Warn("Warn")
	logrus.Error("Error")

	logrus.Info("------------")

	logrus.AddHook(&PrefixHook{Prefix: "[TEST]"})
	logrus.Trace("Trace")
	logrus.Debug("Debug")
	logrus.Info("Info")
	logrus.Warn("Warn")
	logrus.Error("Error")
}
