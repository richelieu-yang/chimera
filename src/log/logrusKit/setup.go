package logrusKit

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var setupOnce sync.Once

func SetUp(config *Config) {
	setupOnce.Do(func() {
		if config == nil {
			config = &Config{
				Level:      "debug",
				PrintBasic: true,
			}
		}

		/* formatter */
		formatter := NewTextFormatter("")
		logrus.SetFormatter(formatter)

		/* level */
		level := StringToLevel(config.Level)
		logrus.SetLevel(level)

		if config.PrintBasic {
			PrintBasicDetails()
		}
	})
	return
}
