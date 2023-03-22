package pulsarKit

import (
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

func NewConsumerOriginally(addresses []string, options pulsar.ConsumerOptions, logPath string) (*Consumer, error) {
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
