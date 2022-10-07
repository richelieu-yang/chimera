package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	var ctx = context.Background()

	// redis://<user>:<pass>@localhost:6379/<db>
	opt, err := redis.ParseURL("redis://:@localhost:6379/0")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	value, err := client.Get(ctx, "test").Result()
	fmt.Println(value)
	fmt.Println(err)
}
