package main

import (
	"github.com/redis/go-redis/v9"
	"github.com/richelieu42/chimera/src/assertKit"
)

func main() {
	assertKit.Must(redis.Nil)
}
