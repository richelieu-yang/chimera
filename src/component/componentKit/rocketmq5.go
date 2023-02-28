package componentKit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/mq/rocketmq5Kit"
	"github.com/sirupsen/logrus"
)

func InitializeRocketMQ5Component() error {
	config, err := GetRocketmq5Config()
	if err != nil {
		return err
	}
	if config == nil {
		return errorKit.Simple("config == nil")
	}

	if err := rocketmq5Kit.VerifyEndpoint(config.Endpoint, config.TopicToVerify); err != nil {
		return errorKit.Wrap(err, "fail to verify endpoint(%s) with topic(%s)", config.Endpoint, config.TopicToVerify)
	}

	logrus.Info("[COMPONENT, ROCKETMQ5] Initialize successfully.")
	return nil
}

// NewProducerOfRocketmq5
/*
!!!: 返回的 rmq_client.Producer 实例要手动调用 Start()（error为nil的情况）.

@param logConfig 可以为nil
*/
func NewProducerOfRocketmq5(logConfig *rocketmq5Kit.LogConfig) (rmq_client.Producer, error) {
	config, err := GetRocketmq5Config()
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, errorKit.Simple("config == nil")
	}

	return rocketmq5Kit.NewProducer(logConfig, &config.Config)
}

// NewSimpleConsumerOfRocketmq5
/*
!!!: 返回的 rmq_client.SimpleConsumer 实例要手动调用 Start()（error为nil的情况）.

@param logConfig 可以为nil
*/
func NewSimpleConsumerOfRocketmq5(logConfig *rocketmq5Kit.LogConfig, consumerGroup, topic, tag string) (rmq_client.SimpleConsumer, error) {
	config, err := GetRocketmq5Config()
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, errorKit.Simple("config == nil")
	}
	return rocketmq5Kit.NewSimpleConsumer(logConfig, &config.Config, consumerGroup, topic, tag)
}
