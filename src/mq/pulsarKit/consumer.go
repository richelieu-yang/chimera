package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/core/errorKit"
)

type (
	Consumer struct {
		pulsar.Client
		pulsar.Consumer
	}
)

func (c *Consumer) Close() {
	if c.Consumer != nil {
		c.Consumer.Close()
	}
	if c.Client != nil {
		c.Client.Close()
	}
}

// NewConsumerOriginally
/*
PS: 目标Pulsar服务未启动的情况下，如果ctx不加以限制，要过约 1min 才会返回error（期间客户端日志有connection refused输出）.
*/
func NewConsumerOriginally(ctx context.Context, addresses []string, options pulsar.ConsumerOptions, logPath string) (*Consumer, error) {
	client, err := NewClient(addresses, logPath)
	if err != nil {
		return nil, errorKit.Wrap(err, "fail to new client")
	}

	consumer, err := client.Subscribe(options)
	if err != nil {
		return nil, errorKit.Wrap(err, "fail to subscribe")
	}

	return &Consumer{
		Client:   client,
		Consumer: consumer,
	}, nil
}
