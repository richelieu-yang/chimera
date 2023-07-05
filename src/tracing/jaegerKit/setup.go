package jaegerKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
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
	if err := interfaceKit.AssertNotNil(config, "config"); err != nil {
		return err
	}

	if !config.Access {
		// case: 不使用jaeger服务
		return nil
	}

	// TODO:

	return nil
}
