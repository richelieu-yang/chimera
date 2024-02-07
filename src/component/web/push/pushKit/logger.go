package pushKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger = logrus.StandardLogger()

func SetDefaultLogger(logrusLogger *logrus.Logger) error {
	if err := interfaceKit.AssertNotNil(logrusLogger, "logrusLogger"); err != nil {
		return err
	}

	logger = logrusLogger
	return nil
}

func GetDefaultLogger() *logrus.Logger {
	return logger
}
