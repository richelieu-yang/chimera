package mqKit

import (
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/sirupsen/logrus"
)

var config *Config

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

	s := sliceKit.RemoveEmpty(config.Endpoints, true)
	if sliceKit.IsEmpty(s) {
		return errorKit.New("config.Endpoint is empty")
	}

	if config.Credentials == nil {
		config.Credentials = &credentials.SessionCredentials{
			AccessKey:    "",
			AccessSecret: "",
		}
	}

	return nil
}
