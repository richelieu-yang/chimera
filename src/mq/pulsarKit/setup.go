package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/assertKit"
	"github.com/richelieu42/chimera/src/core/errorKit"
)

var client pulsar.Client

func MustSetUp(config *Config) {
	err := SetUp(config)
	assertKit.Must(err)
}

func SetUp(config *Config) error {
	var err error

	client, err = NewClient(config)
	if err != nil {
		return errorKit.Wrap(err, "fail to new pulsar client")
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
