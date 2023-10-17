package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"github.com/sirupsen/logrus"
)

var config *Config

// MustSetUp
/*
@param topicForVerify 用于验证的topic（为""则不验证）
*/
func MustSetUp(config *Config, topicForVerify string) {
	if err := SetUp(config, topicForVerify); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(pc *Config, topicForVerify string) (err error) {
	defer func() {
		if err != nil {
			config = nil
		}
	}()

	v := validateKit.New()
	if err = v.Struct(pc); err != nil {
		return
	}
	config = pc

	// verify
	if err = verify(topicForVerify); err != nil {
		return
	}

	return
}

// NewProducer
/*
前提: 成功调用 SetUp() || MustSetUp().

@param options 必需属性: Topic、SendTimeout
@param logPath 客户端的日志输出（"": 输出到控制台）
*/
func NewProducer(ctx context.Context, options pulsar.ProducerOptions, clientLogPath string) (*Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	return NewProducerOriginally(ctx, config.Addrs, options, clientLogPath)
}

// NewConsumer
/*
前提: 成功调用 SetUp() || MustSetUp().

@param options 必需属性: Topic、SubscriptionName、Type
@param logPath 客户端的日志输出（"": 输出到控制台）
*/
func NewConsumer(ctx context.Context, options pulsar.ConsumerOptions, clientLogPath string) (*Consumer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	return NewConsumerOriginally(ctx, config.Addrs, options, clientLogPath)
}
