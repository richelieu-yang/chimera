package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/sirupsen/logrus"
	"sync"
)

var setupOnce sync.Once
var config *Config

func MustSetUp(config *Config) {
	err := SetUp(config)
	if err != nil {
		logrus.Fatal(err)
	}
}

func SetUp(pulsarConfig *Config) (err error) {
	setupOnce.Do(func() {
		config = pulsarConfig
		if config == nil {
			err = errorKit.New("config == nil")
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
前提: 成功调用 SetUp() || MustSetUp().

@param options 必需属性: Topic、SendTimeout
@param logPath 客户端的日志输出（"": 输出到控制台）
*/
func NewProducer(ctx context.Context, options pulsar.ProducerOptions, logPath string) (*Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}
	return NewProducerOriginally(ctx, config.Addresses, options, logPath)
}

// NewConsumer
/*
前提: 成功调用 SetUp() || MustSetUp().

@param options 必需属性: Topic、SubscriptionName、Type
@param logPath 客户端的日志输出（"": 输出到控制台）
*/
func NewConsumer(ctx context.Context, options pulsar.ConsumerOptions, logPath string) (*Consumer, error) {
	if config == nil {
		return nil, NotSetupError
	}
	return NewConsumerOriginally(ctx, config.Addresses, options, logPath)
}
