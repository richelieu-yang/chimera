package operation

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/database/redisKit/test"
	"testing"
)

func TestExists(t *testing.T) {
	fmt.Println("测试 Exists -------------------")

	// client
	client, err := test.NewSingleNodeClient()
	if err != nil {
		panic(err)
	}

	fmt.Println(client.Exists("0", "1", "2"))
	fmt.Println("测试 Exists -------------------")
}
