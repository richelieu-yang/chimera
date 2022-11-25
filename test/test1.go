package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"os"
	"strconv"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang"
)

const (
	Topic     = "wmq"
	Endpoint  = "198.18.0.1:9876"
	AccessKey = "xxxxxx"
	SecretKey = "xxxxxx"
)

func main() {
	if err := os.Setenv("mq.consoleAppender.enabled", "true"); err != nil {
		panic(err)
	}

	rmq_client.ResetLogger()

	// new producer instance
	producer, err := rmq_client.NewProducer(&rmq_client.Config{
		Endpoint:    Endpoint,
		Credentials: &credentials.SessionCredentials{
			//AccessKey:    AccessKey,
			//AccessSecret: SecretKey,
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
		resp, err := producer.Send(context.TODO(), msg)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(resp); i++ {
			fmt.Printf("%#v\n", resp[i])
		}
		// wait a moment
		time.Sleep(time.Second * 1)
	}
}
