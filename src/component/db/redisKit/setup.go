package redisKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var client *Client

func MustSetUp(config *Config) {
	if err := SetUp(config); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(config *Config) (err error) {
	if err := interfaceKit.AssertNotNil(config, "config"); err != nil {
		return err
	}
	client, err = NewClient(*config)
	return
}

// GetClient
/*
前提: 成功调用 SetUp() || MustSetUp().
*/
func GetClient() (*Client, error) {
	if client == nil {
		return nil, NotSetupError
	}

	return client, nil
}
