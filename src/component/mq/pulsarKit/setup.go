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
@param verifyConfig		可以为nil（将不验证，但不推荐这么干）
*/
func MustSetUp(pulsarConfig *Config, verifyConfig *VerifyConfig) {
	if err := SetUp(pulsarConfig, verifyConfig); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(pulsarConfig *Config, verifyConfig *VerifyConfig) (err error) {
	defer func() {
		if err != nil {
			config = nil
		}
	}()

	if err = validateKit.Struct(pulsarConfig); err != nil {
		err = errorKit.Wrap(err, "Fail to verify")
		return
	}
	config = pulsarConfig

	// verify
	if err = verify(verifyConfig); err != nil {
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
