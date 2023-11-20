package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger = logrus.StandardLogger()

func SetDefaultLogger(l *logrus.Logger) error {
	if err := interfaceKit.AssertNotNil(l, "l"); err != nil {
		return err
	}

	logger = l
	return nil
}
