package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
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

func NewProducerOriginally(addresses []string, logPath string, options pulsar.ProducerOptions) (*Producer, error) {
	client, err := NewClient(addresses, logPath)
	if err != nil {
		return nil, err
	}

	producer, err := client.CreateProducer(options)
	if err != nil {
		return nil, err
	}

	return &Producer{
		Client:   client,
		Producer: producer,
	}, nil
}
