package mqKit

import (
	"context"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"time"
)

// NewProducer
/*
PS: In most case, you don't need to create many producers, singletion pattern is more recommended.

@param clientLogPath 客户端日志（blank则输出到控制台）
*/
func NewProducer() (rmq_client.Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	endpoint := sliceKit.Join(config.Endpoints, ";")

	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint:    endpoint,
		Credentials: config.Credentials,
	})
	if err != nil {
		return nil, err
	}
	// Start
	if err := producer.Start(); err != nil {
		return nil, err
	}
	return producer, nil
}

// SendMessage Deprecated: 仅供参考如何生产消息
func SendMessage(producer rmq_client.Producer, ctx context.Context, topic string, tag *string, body []byte, messageGroup string, deliveryTimestamp time.Time, keys ...string) ([]*rmq_client.SendReceipt, error) {
	msg := &rmq_client.Message{
		Topic: topic,
		Tag:   tag,
		Body:  body,
	}
	msg.SetKeys(keys...)
	msg.SetMessageGroup(messageGroup)
	msg.SetDelayTimestamp(deliveryTimestamp)
	return producer.Send(ctx, msg)
}
