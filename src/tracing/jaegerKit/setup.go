package jaegerKit

import (
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func MustSetup(config *Config) {
	if err := Setup(config); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func Setup(config *Config) error {
	// TODO:

	return nil
}
