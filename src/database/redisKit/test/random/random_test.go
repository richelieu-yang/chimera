package random

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/richelieu42/go-scales/src/database/redisKit/test"
	"testing"
)

func TestRandomKey(t *testing.T) {
	fmt.Println("测试 Random ---------------------------------------")

	// client
	client, err := test.NewSingleNodeClient()
	if err != nil {
		panic(err)
	}

	key, err := client.RandomKey()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("当前db为空！")
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("key: [%s].\n", key)
	}

	fmt.Println("测试 Random ---------------------------------------")
}

func Test(t *testing.T) {
	// client
	client, err := test.NewSingleNodeClient()
	if err != nil {
		panic(err)
	}

	fmt.Println(client.Scan(0, "cyy:*", 3))
	fmt.Println(client.Scan(1, "cyy:*", 3))
	fmt.Println(client.Scan(2, "cyy:*", 3))
	fmt.Println(client.Scan(3, "cyy:*", 3))
	fmt.Println(client.Scan(4, "cyy:*", 3))
	fmt.Println(client.Scan(5, "cyy:*", 3))
}
