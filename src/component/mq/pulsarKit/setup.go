package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"github.com/sirupsen/logrus"
)

var config *Config

// MustSetUp
/*
@param topicForVerify 	用于验证的topic（为""则不验证）
@param printFlag 		是否输出 verify相关的输出？
*/
func MustSetUp(config *Config, topicForVerify string, printFlag bool) {
	if err := SetUp(config, topicForVerify, printFlag); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(pc *Config, topicForVerify string, printFlag bool) (err error) {
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
	if err = verify(topicForVerify, printFlag); err != nil {
		err = errorKit.Wrap(err, "Fail to verify")
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
