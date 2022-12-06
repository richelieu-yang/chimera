package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
)

const (
	Topic         = "ttt"
	ConsumerGroup = "cg"
	Endpoint      = "localhost:8081"

	AccessKey = ""
	SecretKey = ""
)

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
)

func main() {
	// log to console
	//os.Setenv("mq.consoleAppender.enabled", "true")

	rmq_client.ResetLogger()
	// new simpleConsumer instance
	simpleConsumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithAwaitDuration(awaitDuration),
		rmq_client.WithSubscriptionExpressions(map[string]*rmq_client.FilterExpression{
			Topic: rmq_client.SUB_ALL,
		}),
	)
	if err != nil {
		panic(err)
	}
	// start simpleConsumer
	err = simpleConsumer.Start()
	if err != nil {
		panic(err)
	}
	// gracefule stop simpleConsumer
	defer simpleConsumer.GracefulStop()

	go func() {
		for {
			mvs, err := simpleConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
			if err != nil {
				// ！！！：此处不能用 panic()
				logrus.Errorf("[CONSUMER] fail to receive, error: %+v", err)
			}

			// ack message
			for _, mv := range mvs {
				if err := simpleConsumer.Ack(context.TODO(), mv); err != nil {
					logrus.Errorf("[CONSUMER] fail to ack, error: %+v", err)
				} else {
					logrus.Infof("[CONSUMER] recevie message(id: %s, text: %s).", mv.GetMessageId(), string(mv.GetBody()))
				}
			}

			logrus.Info("---------------------------------------------------------")

			// wait a moment
			time.Sleep(time.Second * 3)
		}
	}()

	// run for a while
	time.Sleep(time.Minute * 60)
}
