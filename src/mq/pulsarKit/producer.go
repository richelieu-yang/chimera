package pulsarKit

import (
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
	if p.Producer != nil {
		p.Producer.Close()
	}
	if p.Client != nil {
		p.Client.Close()
	}
}

func NewProducerOriginally(addresses []string, options pulsar.ProducerOptions, logPath string) (*Producer, error) {
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
