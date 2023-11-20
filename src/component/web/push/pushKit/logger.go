package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger = logrus.StandardLogger()

func SetDefaultLogger(logger *logrus.Logger) error {
	if err := interfaceKit.AssertNotNil(logger, "logger"); err != nil {
		return err
	}

	logger = logger
	return nil
}
