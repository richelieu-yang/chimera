package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/core/errorKit"
)

type (
	Producer struct {
		pulsar.Client
		pulsar.Producer
	}
)

func (p *Producer) Close() {
	p.Producer.Close()
	p.Client.Close()
}

// NewProducerOriginally
/*
PS: 目标Pulsar服务未启动的情况下，如果ctx不加以限制，要过约 1min 才会返回error（期间客户端日志有connection refused输出）.
*/
func NewProducerOriginally(ctx context.Context, addresses []string, options pulsar.ProducerOptions, logPath string) (*Producer, error) {
	client, err := NewClient(addresses, logPath)
	if err != nil {
		return nil, errorKit.Wrap(err, "fail to new client")
	}

	producer, err := client.CreateProducer(options)
	if err != nil {
		return nil, errorKit.Wrap(err, "fail to create producer")
	}

	return &Producer{
		Client:   client,
		Producer: producer,
	}, nil
}
