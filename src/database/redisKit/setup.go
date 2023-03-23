package redisKit

import (
	"github.com/richelieu42/chimera/src/assertKit"
	"github.com/zeromicro/go-zero/core/logx"
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
	})

	logx.SetUp()

	return
}
