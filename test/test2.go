package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu42/chimera/src/assertKit"
)

func main() {
	err := assertKit.Nil(redis.Nil, "name")
	fmt.Println(err.Error())
}
