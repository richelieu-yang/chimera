package main

import (
	"context"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang"
)

const (
	Topic = "test123"
	// Endpoint Proxy服务的ip+port
	Endpoint  = "localhost:8081"
	AccessKey = ""
	SecretKey = ""
)

func main() {
	if err := os.Setenv("mq.consoleAppender.enabled", "true"); err != nil {
		panic(err)
	}

	rmq_client.ResetLogger()

	// new producer instance
	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint: Endpoint,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithTopics(Topic),
	)
	if err != nil {
		panic(err)
	}

	// start producer
	if err := producer.Start(); err != nil {
		panic(err)
	}

	// gracefule stop producer
	defer producer.GracefulStop()

	for i := 0; i < 10; i++ {
		// new a message
		msg := &rmq_client.Message{
			Topic: Topic,
			Body:  []byte("this is a message : " + strconv.Itoa(i)),
		}
		// set keys and tag
		msg.SetKeys("a", "b")
		msg.SetTag("ab")

		// send message in sync
		resps, err := producer.Send(context.TODO(), msg)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(resps); i++ {
			resp := resps[i]
			logrus.Infof("MessageID: [%s].", resp.MessageID)
		}

		// wait a moment
		time.Sleep(time.Second * 1)
	}
}
