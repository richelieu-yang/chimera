package redisKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
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
	if err = validateKit.Struct(config); err != nil {
		err = errorKit.Wrap(err, "Fail to verify")
		return
	}

	client, err = NewClient(*config)
	return
}

// IsAccess
/*
@return true: 已经正确setup了
*/
func IsAccess() bool {
	return client != nil
}

func MustGetClient() *Client {
	client, err := GetClient()
	if err != nil {
		logrus.Fatal(err)
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
