package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/idKit"
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewConsumerOriginally(t *testing.T) {
	address := []string{"192.168.80.27:6650", "192.168.80.42:6650", "192.168.80.43:6650"}
	topic := "test"
	logPath := "logs/pulsar-consumer.log"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()

	consumer, err := NewConsumerOriginally(ctx, address, pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: idKit.NewULID(),
		Type:             pulsar.Exclusive,
	}, logPath)
	assert.Nil(t, err)

	logrusKit.SetUp(&logrusKit.Config{
		Level:      "debug",
		PrintBasic: false,
	})
	for cMsg := range consumer.Chan() {
		if err := cMsg.Ack(cMsg); err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("[CONSUMER] fail to ack")
			continue
		}
		logrus.WithFields(logrus.Fields{
			"text": string(cMsg.Payload()),
		}).Info("[CONSUMER] receive a message")
	}
}
