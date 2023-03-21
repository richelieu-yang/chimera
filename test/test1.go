package main

import (
	"github.com/redis/go-redis/v9"
	"github.com/richelieu42/chimera/src/assertKit"
	"github.com/richelieu42/chimera/src/core/errorKit"
)

func main() {
	var err error = redis.Nil
	err = errorKit.Wrap(err, "Redis错误")
	assertKit.Must(err)
}
