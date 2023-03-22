package main

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/mq/pulsarKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	addresses := []string{"localhost:6650"}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	consumer, err := pulsarKit.NewConsumerOriginally(ctx, addresses, pulsar.ConsumerOptions{
		Topic:            "test",
		SubscriptionName: "my-sub1",
		Type:             pulsar.Exclusive,
	}, "")
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(consumer)
}
