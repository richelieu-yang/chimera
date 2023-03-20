package pulsarKit

import "github.com/apache/pulsar-client-go/pulsar"

func NewClient() (pulsar.Client, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
}
