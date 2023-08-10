package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.TODO())

	cancel(redis.Nil)
	fmt.Println(ctx.Err())                       // context canceled
	fmt.Println(ctx.Err() == context.Canceled)   // true
	fmt.Println(context.Cause(ctx))              // redis: nil
	fmt.Println(context.Cause(ctx) == redis.Nil) // true
}
