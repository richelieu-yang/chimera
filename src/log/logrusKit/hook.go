package logrusKit

import (
	"github.com/sirupsen/logrus"
)

// 保护性检查
var (
	_ logrus.Hook = (*defaultPrefixHook)(nil)
)

type defaultPrefixHook struct {
	prefix string
}

func (hook *defaultPrefixHook) Fire(entry *logrus.Entry) error {
	//if strKit.IsNotEmpty(hook.prefix) {
	//	entry.Message = fmt.Sprintf("%s %s", hook.prefix, entry.Message)
	//}
	entry.Message = hook.prefix + entry.Message

	return nil
}

func (hook *defaultPrefixHook) Levels() []logrus.Level {
	//// 只有 INFO、WARN 级别
	//return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel}

	return logrus.AllLevels
}
