package main

import (
	"context"
	"fmt"
	"github.com/richelieu42/go-scales/src/database/redisKit"
)

func main() {
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
