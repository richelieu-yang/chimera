package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu42/chimera/src/core/errorKit"
)

func main() {
	var err error = redis.Nil
	err = errorKit.WithMessage(err, "message")

	//test/test2.go:11|main message: redis: nil
	fmt.Printf("%v\n", err)

	//redis: nil
	//test/test2.go:11|main message
	fmt.Printf("%+v\n", err)
}
