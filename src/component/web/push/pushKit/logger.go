package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/sirupsen/logrus"
)

var defLogger *logrus.Logger = logrus.StandardLogger()

func SetDefaultLogger(logger *logrus.Logger) error {
	if err := interfaceKit.AssertNotNil(logger, "logger"); err != nil {
		return err
	}

	defLogger = logger
	return nil
}
