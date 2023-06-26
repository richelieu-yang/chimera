package main

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/mq/pulsarKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(nil)

	pulsarKit.MustSetUp(&pulsarKit.Config{
		Addresses: []string{"pulsar://localhost:6650"},
		VerifyConfig: &pulsarKit.VerifyConfig{
			Topic: "test",
			Print: true,
		},
	})

	consumer, err := pulsarKit.NewConsumer(context.TODO(), pulsar.ConsumerOptions{
		Topic:            "test",
		SubscriptionName: "name",
		Type:             pulsar.Exclusive,
	}, "")
	if err != nil {
		logrus.Fatal(err)
	}
	consumer.Close()

	consumer1, err := pulsarKit.NewConsumer(context.TODO(), pulsar.ConsumerOptions{
		Topic:            "test",
		SubscriptionName: "name",
		Type:             pulsar.Exclusive,
	}, "")
	if err != nil {
		logrus.Fatal(err)
	}
	defer consumer1.Close()

	select {}

}
