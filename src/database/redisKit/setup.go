package redisKit

import (
	"github.com/richelieu42/chimera/src/assertKit"
	"sync"
)

var client *Client
var setupOnce sync.Once

func MustSetUp(config *Config) {
	err := SetUp(config)
	assertKit.Must(err)
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
