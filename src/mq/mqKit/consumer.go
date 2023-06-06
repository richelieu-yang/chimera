package mqKit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// NewSimpleConsumer
/*
PS: In most case, you don't need to create many consumers, singletion pattern is more recommended.

@param consumerGroup 			不能为""
@param subscriptionExpressions	(1) key: 	topic，不能为 "*" || blank
								(2) value: 	tag，一般为 rmq_client.SUB_ALL
@param clientLogPath 			客户端日志（blank则输出到控制台）
*/
func NewSimpleConsumer(consumerGroup string, subscriptionExpressions map[string]*rmq_client.FilterExpression) (rmq_client.SimpleConsumer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	endpoint := sliceKit.Join(config.Endpoints, ";")
	if err := strKit.AssertNotEmpty(consumerGroup, "consumerGroup"); err != nil {
		return nil, err
	}
	//if err := mapKit.AssertNotEmpty(subscriptionExpressions, "subscriptionExpressions"); err != nil {
	//	return nil, err
	//}

	simpleConsumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:      endpoint,
		ConsumerGroup: consumerGroup,
		Credentials:   config.Credentials,
	},
		rmq_client.WithAwaitDuration(AwaitDuration),
		rmq_client.WithSubscriptionExpressions(subscriptionExpressions),
	)
	if err != nil {
		return nil, err
	}
	// Start
	if err := simpleConsumer.Start(); err != nil {
		return nil, err
	}
	return simpleConsumer, nil
}
