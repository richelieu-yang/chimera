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
@param config 可以为nil
*/
func SetUp(config *Config) error {
	if config == nil {
		config = &Config{
			Level:      "debug",
			PrintBasic: false,
		}
	}

	logrus.SetFormatter(DefaultTextFormatter)
	logrus.SetReportCaller(true)
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
