package redisKit

import (
	"github.com/sirupsen/logrus"
	"sync"
)

var setupOnce sync.Once
var client *Client

func MustSetUp(config *Config) {
	err := SetUp(config)
	if err != nil {
		logrus.Fatal(err)
	}
}

func SetUp(config *Config) (err error) {
	setupOnce.Do(func() {
		client, err = NewClient(config)

		if err != nil {
			client = nil
		}
	})

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
