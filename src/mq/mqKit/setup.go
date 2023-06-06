package mqKit

import (
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/sirupsen/logrus"
)

func MustSetUp(config *Config) {
	if err := SetUp(config); err != nil {
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(config *Config) error {
	if err := check(config); err != nil {
		return err
	}

	return nil
}

func check(config *Config) error {
	if config == nil {
		return errorKit.New("config == nil")
	}

	if strKit.IsBlank(config.Endpoint) {
		return errorKit.New("config.Endpoint is blank")
	}
	if config.Credentials == nil {
		config.Credentials = &credentials.SessionCredentials{
			AccessKey:    "",
			AccessSecret: "",
		}
	}

	return nil
}
