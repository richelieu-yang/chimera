package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/assertKit"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"sync"
)

var client pulsar.Client
var setupOnce sync.Once

func MustSetUp(config *Config) {
	err := SetUp(config)
	assertKit.Must(err)
}

func SetUp(config *Config) error {
	var err error

	setupOnce.Do(func() {
		client, err = NewClient(config)
	})

	if err != nil {
		return errorKit.Wrap(err, "fail to set up")
	}
	return nil
}

// GetClient
/*
PS: 必须先成功调用 SetUp || MustSetUp.
*/
func GetClient() (pulsar.Client, error) {
	if client == nil {
		return nil, NotSetupError
	}
	return client, nil
}
