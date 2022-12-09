package rocketmq5Kit

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

var (
	topic = "test"

	config = &rmq_client.Config{
		Endpoint:    "localhost:8081",
		Credentials: nil,
	}

	producerLogConfig = &LogConfig{
		ToConsole: false,
		LogDir:    "",
		LogName:   "producer.log",
	}

	consumerLogConfig = &LogConfig{
		ToConsole: false,
		LogDir:    "",
		LogName:   "consumer.log",
	}
)

// 用于测试：RocketMQ5服务是否启动成功？
func TestConnectToRocketMQ5(t *testing.T) {
	go func() {
		time.Sleep(time.Second * 3)

		if err := startProducer(); err != nil {
			logrus.Panic(err)
		}
	}()

	if err := startConsumer(); err != nil {
		logrus.Panic(err)
	}
}

func startProducer() error {
	producer, err := NewProducer(producerLogConfig, config)
	if err != nil {
		return err
	}
	if err := producer.Start(); err != nil {
		return err
	}

	defer producer.GracefulStop()

	for i := 0; i < 3; i++ {
		text := fmt.Sprintf("%d %s", i, timeKit.FormatCurrentTime())
		receipt, err := SendMessage(producer, topic, []byte(text), nil)
		if err != nil {
			return err
		}
		logrus.Infof("[Producer] MessageID: [%s], Offset: [%d].", receipt.MessageID, receipt.Offset)
	}
	logrus.Info("[Producer] End.")
	return nil
}

func startConsumer() error {
	simpleConsumer, err := NewSimpleConsumer(consumerLogConfig, config, "cg", topic, "*")
	if err != nil {
		return err
	}
	if err := simpleConsumer.Start(); err != nil {
		return err
	}

	defer simpleConsumer.GracefulStop()

	go func() {
		for {
			mvs, err := simpleConsumer.Receive(context.TODO(), MaxMessageNum, InvisibleDuration)
			if err != nil {
				// ！！！：此处不能用 panic()
				logrus.Errorf("[CONSUMER] fail to receive, error: %+v", err)
			}

			// ack message
			for _, mv := range mvs {
				if err := simpleConsumer.Ack(context.TODO(), mv); err != nil {
					logrus.Errorf("[CONSUMER] fail to ack, error: %+v", err)
				} else {
					logrus.Infof("[CONSUMER] recevie message(id: %s, text: %s, tag: %v, offset: %d).", mv.GetMessageId(), string(mv.GetBody()), mv.GetTag(), mv.GetOffset())
				}
			}

			// wait a moment
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// run for a while
	time.Sleep(time.Second * 10)

	logrus.Info("[CONSUMER] End.")
	return nil
}
