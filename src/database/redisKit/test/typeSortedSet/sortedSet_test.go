package typeSortedSet

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/database/redisKit/test"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestSortedSet(t *testing.T) {
	fmt.Println("测试 sorted set---------------------------------------")

	// client
	client, err := test.NewSingleNodeClient()
	if err != nil {
		panic(err)
	}

	s, err := client.ZRevRangeByScore("zset", &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  3,
	})
	fmt.Println(s, err)

	fmt.Println("测试 sorted set---------------------------------------")
}
