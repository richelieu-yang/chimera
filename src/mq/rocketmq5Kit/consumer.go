package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
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
@param logConfig	可以为nil（此时将默认输出到控制台）
@param config 		不会修改传参config，因为修改的是副本
@return rmq_client.SimpleConsumer实例要手动调用 Start()!!!
*/
func NewSimpleConsumer(logConfig *LogConfig, config *rmq_client.Config, consumerGroup, topic, tag string) (rmq_client.SimpleConsumer, error) {
	lock.Lock()
	defer lock.Unlock()

	if err := logConfig.SetLogout(); err != nil {
		return nil, err
	}

	config, err := processConfig(config)
	if err != nil {
		return nil, err
	}

	if strKit.IsEmpty(consumerGroup) {
		return nil, errorKit.New("consumerGroup is empty")
	}
	if strKit.IsEmpty(topic) {
		return nil, errorKit.New("topic is empty")
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
