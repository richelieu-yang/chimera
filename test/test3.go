package main

import (
	"github.com/gogf/gf/v2/os/gcache"
)

func main() {
	cache := gcache.New()
	cache.Set("k1", "v1", 0)
}
