package pulsarKit

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/v2/core/file/fileKit"
	"github.com/richelieu42/chimera/v2/log/logrusKit"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewProducerOriginally(t *testing.T) {
	address := []string{"192.168.80.27:6650", "192.168.80.42:6650", "192.168.80.43:6650"}
	topic := "test"
	sendTimeout := time.Second * 3

	logPath := "logs/pulsar-producer.log"
	err := fileKit.Delete(logPath)
	assert.Nil(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()

	producer, err := NewProducerOriginally(ctx, address, pulsar.ProducerOptions{
		Topic:       topic,
		SendTimeout: sendTimeout,
	}, logPath)
	assert.Nil(t, err)

	logrusKit.SetUp(&logrusKit.Config{
		Level:      "debug",
		PrintBasic: false,
	})
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second * 3)

		text := fmt.Sprintf("message-%d", i)
		_, err := producer.SendWithTimeout(&pulsar.ProducerMessage{
			Payload: []byte(text),
		}, sendTimeout)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Error("[PRODUCER] fail to send")
		} else {
			logrus.WithFields(logrus.Fields{
				"text": text,
			}).Info("[PRODUCER] manager to send")
		}
	}
}
