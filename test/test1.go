package main

import (
	"context"
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/richelieu42/go-scales/src/mq/rocketmq5Kit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	config := &rmq_client.Config{
		Endpoint:    "localhost:8081",
		Credentials: nil,
	}

	consumer, err := rocketmq5Kit.NewSimpleConsumer(nil, config, "cg", "test", "*")
	if err != nil {
		panic(err)
	}
	if err := consumer.Start(); err != nil {
		panic(err)
	}

	// gracefule stop simpleConsumer
	defer consumer.GracefulStop()

	go func() {
		for {
			mvs, err := consumer.Receive(context.TODO(), rocketmq5Kit.MaxMessageNum, rocketmq5Kit.InvisibleDuration)
			if err != nil {
				// ！！！：此处不能用 panic()
				logrus.Errorf("[CONSUMER] fail to receive, error: %+v", err)
			}

			// ack message
			for _, mv := range mvs {
				if err := consumer.Ack(context.TODO(), mv); err != nil {
					logrus.Errorf("[CONSUMER] fail to ack, error: %+v", err)
				} else {
					logrus.Infof("[CONSUMER] recevie message(id: %s, text: %s).", mv.GetMessageId(), string(mv.GetBody()))
				}
			}

			logrus.Info("---------------------------------------------------------")

			// wait a moment
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// run for a while
	time.Sleep(time.Minute * 60)

}
