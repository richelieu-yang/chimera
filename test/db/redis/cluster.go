package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	var ctx = context.Background()

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"172.18.21.32:8091", "172.18.21.32:8092", "172.18.21.32:8093"},

		IdleTimeout:  3 * time.Minute, // 默认5min
		MinIdleConns: 32,
		Username:     "",
		Password:     "",

		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		//RouteRandomly: true,
	})

	value, err := client.Get(ctx, "test").Result()
	fmt.Println(value)
	fmt.Println(err)
}
