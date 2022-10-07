package componentKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func initializeLogrusComponent() error {
	config, err := GetLogrusConfig()
	if err != nil {
		return err
	}
	if config == nil {
		return errorKit.Simple("config == nil")
	}

	level := logrusKit.StringToLevel(config.Level)
	logrusKit.Initialize(level, timeKit.TimeFormat(config.TimestampFormat))
	logrusKit.PrintBasicDetails()

	logrus.Info("[COMPONENT, LOGRUS] Initialize successfully.")
	return nil
}
