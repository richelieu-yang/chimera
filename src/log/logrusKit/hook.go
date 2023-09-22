package logrusKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/sirupsen/logrus"
)

type defaultPrefixHook struct {
	logrus.Hook

	prefix string
}

func (hook *defaultPrefixHook) Fire(entry *logrus.Entry) error {
	if strKit.IsNotEmpty(hook.prefix) {
		entry.Message = fmt.Sprintf("%s %s", hook.prefix, entry.Message)
	}
	return nil
}

func (hook *defaultPrefixHook) Levels() []logrus.Level {
	//return logrus.AllLevels

	// 只有 INFO、WARN 级别
	return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}
}
