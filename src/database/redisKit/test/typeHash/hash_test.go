package typeHash

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/database/redisKit/test"
	"testing"
)

func TestHash(t *testing.T) {
	fmt.Println("测试 hash ---------------------------------------")

	// client
	client, err := test.NewSingleNodeClient()
	if err != nil {
		panic(err)
	}

	str, err := client.HGet("11", "3")
	fmt.Println(str, err)

	fmt.Println("测试 hash ---------------------------------------")
}
