package main

import (
	"context"
	"fmt"
	"os"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
)

const (
	Topic = "ccc"
	// Endpoint Proxy服务的ip+port
	Endpoint  = "localhost:8081"
	AccessKey = ""
	SecretKey = ""

	ConsumerGroup = "cg"
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
	os.Setenv("mq.consoleAppender.enabled", "true")

	rmq_client.ResetLogger()

	// new consumer instance
	consumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
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
	// start consumer
	err = consumer.Start()
	if err != nil {
		panic(err)
	}
	// gracefule stop consumer
	defer consumer.GracefulStop()

	go func() {
		for {
			fmt.Println("start recevie message")
			mvs, err := consumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
			if err != nil {
				panic(err)
			}
			// ack message
			for _, mv := range mvs {
				consumer.Ack(context.TODO(), mv)
				fmt.Println(mv)
			}
			fmt.Println("wait a moment")
			fmt.Println()
			time.Sleep(time.Second * 3)
		}
	}()
	// run for a while
	time.Sleep(time.Minute)
}
