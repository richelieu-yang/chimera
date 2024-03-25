package redisKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var (
	NotSetupError = errorKit.Newf("haven’t been set up correctly")
)

var client *Client

func MustSetUp(config *Config) {
	if err := SetUp(config); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(config *Config) (err error) {
	client, err = NewClient(config)
	return
}

// GetClientInsecurely
/*
PS: 可能会panic.
*/
func GetClientInsecurely() *Client {
	client, err := GetClient()
	if err != nil {
		logrus.Panic(err)
	}
	return client
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
