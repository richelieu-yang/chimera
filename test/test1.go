package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/mq/rocketmq5Kit"
)

func main() {
	if err := rocketmq5Kit.TestEndpoint("localhost:8081", "test"); err != nil {
		panic(err)
	}
	fmt.Println("---------------------------")
}
