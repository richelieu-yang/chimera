package main

import (
	"context"
	"fmt"
	"github.com/richelieu42/go-scales/src/database/redisKit"
)

func main() {
	//logger, err := logrusKit.NewFileLogger("c/a.log", nil, logrus.DebugLevel, false)
	//if err != nil {
	//	panic(err)
	//}
	//logger.Info("0")
	//logger.Info("1")
	//logger.Info("2")

	config := &redisKit.RedisConfig{
		UserName: "",
		Password: "",
		Mode:     redisKit.SingleNodeMode,
		SingleNodeConfig: &redisKit.SingleNodeConfig{
			Addr: "127.0.0.1:6379",
			DB:   10,
		},
		MasterSlaverConfig: nil,
		SentinelConfig:     nil,
		ClusterConfig:      nil,
	}
	client, err := redisKit.NewClient(config)
	if err != nil {
		panic(err)
	}

	pubsub := client.Subscribe(context.TODO(), "")
	ch := pubsub.Channel()
	//fmt.Println(client.Ping(context.TODO()))
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
