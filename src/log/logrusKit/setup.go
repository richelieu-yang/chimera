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

@param config 可以为nil（debug级别、禁用双引号、不输出基本信息）
*/
func SetUp(config *Config) error {
	if config == nil {
		config = &Config{
			Level:       "debug",
			EnableQuote: false,
			PrintBasic:  false,
		}
	}

	logrus.SetFormatter(NewDefaultTextFormatter())
	logrus.SetReportCaller(true)

	/* Level */
	level, err := ParseLevel(config.Level)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	/* EnableQuote */
	if config.EnableQuote {
		EnableQuote(nil)
	} else {
		DisableQuote(nil)
	}

	/* EnableQuote */
	if config.PrintBasic {
		PrintBasicDetails(nil)
	}
	return nil
}
