package redisKit

import (
	"github.com/sirupsen/logrus"
)

var client *Client

func MustSetUp(config *Config) {
	if err := SetUp(config); err != nil {
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(config *Config) (err error) {
	client, err = NewClient(config)
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
