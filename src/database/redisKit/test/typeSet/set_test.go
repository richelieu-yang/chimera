package typeSet

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/database/redisKit/test"
	"testing"
)

func TestHash(t *testing.T) {
	fmt.Println("测试 hash ---------------------------------------")

	// client
	client, err := test.NewSingleNodeClient()
	if err != nil {
		panic(err)
	}

	fmt.Println(client.SAdd("set", "1", "2", "3", "2", "1"))

	fmt.Println("测试 hash ---------------------------------------")
}
