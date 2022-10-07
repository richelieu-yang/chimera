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

	// 键"ccc"对应的值不存在
	value, err := client.Get(ctx, "ccc").Result()
	if err == redis.Nil {
		fmt.Printf("key(%s) doesn't exist.\n", "ccc")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("value: [%s].\n", value)
	}
}
