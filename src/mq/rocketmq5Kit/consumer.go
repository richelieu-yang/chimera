package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"time"
)

const (
	// AwaitDuration (consumer创建实例时用到) maximum waiting time for receive func
	AwaitDuration = time.Second * 5

	// MaxMessageNum (consumer收消息时用到) maximum number of messages received at one time
	MaxMessageNum int32 = 16

	// InvisibleDuration (consumer收消息时用到) InvisibleDuration should > 20s
	InvisibleDuration = time.Second * 20

	// receive messages in a loop
)

// NewSimpleConsumer
/*
@param logConfig 	可以为nil（此时将默认输出到控制台）
@param config 		不会修改传参config，因为修改的是副本
*/
func NewSimpleConsumer(logConfig *LogConfig, config *rmq_client.Config, consumerGroup, topic, tag string) (rmq_client.SimpleConsumer, error) {
	lock.Lock()
	defer lock.Unlock()

	if err := logConfig.SetLogout(); err != nil {
		return nil, err
	}

	config, err := processBaseConfig(config)
	if err != nil {
		return nil, err
	}

	if strKit.IsEmpty(consumerGroup) {
		return nil, errorKit.Simple("consumerGroup is empty")
	}
	if strKit.IsEmpty(topic) {
		return nil, errorKit.Simple("topic is empty")
	}
	tag = strKit.EmptyToDefault(tag, "*")

	config.ConsumerGroup = consumerGroup

	return rmq_client.NewSimpleConsumer(config,
		rmq_client.WithAwaitDuration(AwaitDuration),
		rmq_client.WithSubscriptionExpressions(map[string]*rmq_client.FilterExpression{
			topic: rmq_client.NewFilterExpression(tag),
		}),
	)
}
