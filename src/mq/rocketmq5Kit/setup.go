package rocketmq5Kit

import (
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var config *Config

func MustSetUp(config *Config) {
	if err := SetUp(config); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(c *Config) error {
	if err := check(c); err != nil {
		return err
	}

	if err := initClientLog(c.ClientLogPath); err != nil {
		return err
	}

	config = c
	return nil
}

func check(c *Config) error {
	if c == nil {
		return errorKit.New("c == nil")
	}

	// Endpoints
	s := sliceKit.RemoveEmpty(c.Endpoints, true)
	if err := sliceKit.AssertNotEmpty(s, "c.Endpoints"); err != nil {
		return err
	}
	c.Endpoints = s

	// Credentials
	if c.Credentials == nil {
		c.Credentials = &credentials.SessionCredentials{
			AccessKey:     "",
			AccessSecret:  "",
			SecurityToken: "",
		}
	}

	return nil
}
