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
	err := SetUp(config)
	assertKit.Must(err)
}

func SetUp(pulsarConfig *Config) (err error) {
	setupOnce.Do(func() {
		config = pulsarConfig

		if config == nil {
			err = errorKit.Simple("pulsarConfig == nil")
		} else {
			err = verify(config.VerifyConfig)
		}

		if err != nil {
			config = nil
		}
	})

	return err
}

// NewProducer
/*
@param logPath 客户端的日志输出（为空则输出到控制台）
*/
func NewProducer(ctx context.Context, options pulsar.ProducerOptions, logPath string) (*Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}
	return NewProducerOriginally(ctx, config.Addresses, options, logPath)
}

// NewConsumer
/*
@param options 至少需要为 Topic、SubscriptionName、Type 属性复制
@param logPath 客户端的日志输出（为空则输出到控制台）
*/
func NewConsumer(ctx context.Context, options pulsar.ConsumerOptions, logPath string) (*Consumer, error) {
	if config == nil {
		return nil, NotSetupError
	}
	return NewConsumerOriginally(ctx, config.Addresses, options, logPath)
}
