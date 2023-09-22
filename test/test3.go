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
	return logrus.AllLevels
}

func main() {
	logrus.Info("0") // time="2023-09-22T14:37:33+08:00" level=info msg=0

	logrus.AddHook(&PrefixHook{Prefix: "[TEST]"})
	logrus.Info("1") // time="2023-09-22T14:37:33+08:00" level=info msg="[TEST] 1"
}
