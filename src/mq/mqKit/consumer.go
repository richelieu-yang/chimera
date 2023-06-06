package mqKit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// NewSimpleConsumer
/*
PS:
(1) In most case, you don't need to create many consumers, singletion pattern is more recommended.
*/
func NewSimpleConsumer(consumerGroup string, subscriptionExpressions map[string]*rmq_client.FilterExpression) (rmq_client.SimpleConsumer, error) {
	endpoint := sliceKit.Join(config.Endpoints, ";")
	if err := strKit.AssertStringNotEmpty(consumerGroup); err != nil {
		return nil, err
	}

	return rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:      endpoint,
		ConsumerGroup: consumerGroup,
		Credentials:   config.Credentials,
	},
		rmq_client.WithAwaitDuration(AwaitDuration),
		rmq_client.WithSubscriptionExpressions(subscriptionExpressions),
	)
}
