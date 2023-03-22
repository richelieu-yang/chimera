package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/assertKit"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"sync"
)

var setupOnce sync.Once
var config *Config

func MustSetUp(config *Config) {
	assertKit.Must(SetUp(config))
}

func SetUp(pulsarConfig *Config) (err error) {
	if pulsarConfig == nil {
		return errorKit.Simple("pulsarConfig == nil")
	}

	setupOnce.Do(func() {
		config = pulsarConfig
		err = verify(config.VerifyConfig)
	})

	return err
}

func NewProducer(options pulsar.ProducerOptions, logPath string) (*Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}
	return NewProducerOriginally(config.Addresses, options, logPath)
}

func NewConsumer(options pulsar.ConsumerOptions, logPath string) (*Consumer, error) {
	if config == nil {
		return nil, NotSetupError
	}
	return NewConsumerOriginally(config.Addresses, options, logPath)
}
