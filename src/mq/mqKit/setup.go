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

	if err := initClientLog(config.ClientLogPath); err != nil {
		return err
	}

	return nil
}

func check(config *Config) error {
	if config == nil {
		return errorKit.New("config == nil")
	}

	// Endpoints
	s := sliceKit.RemoveEmpty(config.Endpoints, true)
	if err := sliceKit.AssertNotEmpty(s, "config.Endpoints"); err != nil {
		return err
	}
	config.Endpoints = s

	// Credentials
	if config.Credentials == nil {
		config.Credentials = &credentials.SessionCredentials{
			AccessKey:    "",
			AccessSecret: "",
		}
	}

	return nil
}
