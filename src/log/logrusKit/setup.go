package logrusKit

import (
	"github.com/sirupsen/logrus"
)

func MustSetUp(config *Config) {
	err := SetUp(config)
	if err != nil {
		DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// SetUp
/*
PS: 可以多次调用.

@param config 可以为nil（采用默认值）
*/
func SetUp(config *Config) error {
	if config == nil {
		config = &Config{
			Level:        "debug",
			PrintBasic:   false,
			DisableQuote: false,
		}
	}

	logrus.SetFormatter(NewDefaultTextFormatter())
	logrus.SetReportCaller(true)

	// 默认日志级别: Debug
	level, err := ParseLevel(config.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	if config.PrintBasic {
		PrintBasicDetails(nil)
	}
	return nil
}
