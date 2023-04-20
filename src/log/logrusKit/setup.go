package logrusKit

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var setupOnce sync.Once

func MustSetUp(config *Config) {
	err := SetUp(config)
	if err != nil {
		logrus.Fatal(err)
	}
}

// SetUp
/*
@param config 可以为nil
*/
func SetUp(config *Config) (err error) {
	setupOnce.Do(func() {
		if config == nil {
			config = &Config{
				Level:      "debug",
				PrintBasic: false,
			}
		}

		logrus.SetFormatter(DefaultTextFormatter)
		logrus.SetReportCaller(true)
		var level logrus.Level
		level, err = ParseLevel(config.Level)
		if err != nil {
			return
		}
		logrus.SetLevel(level)

		if config.PrintBasic {
			PrintBasicDetails()
		}
	})
	return
}
