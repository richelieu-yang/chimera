package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	var ctx = context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		DB:           0,               // use default DB
		IdleTimeout:  3 * time.Minute, // 默认5min
		MinIdleConns: 32,
		Username:     "",
		Password:     "", // no password set
	})

	value, err := client.Ping(ctx).Result()
	fmt.Println(value, err) // PONG <nil>

	value, err = client.Get(ctx, "test").Result()
	fmt.Println(value, err)
}
