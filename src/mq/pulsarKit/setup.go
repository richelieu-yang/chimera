package pulsarKit

import (
	"context"
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

func NewProducer(ctx context.Context, options pulsar.ProducerOptions, logPath string) (*Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}
	return NewProducerOriginally(ctx, config.Addresses, options, logPath)
}

func NewConsumer(ctx context.Context, options pulsar.ConsumerOptions, logPath string) (*Consumer, error) {
	if config == nil {
		return nil, NotSetupError
	}
	return NewConsumerOriginally(ctx, config.Addresses, options, logPath)
}
