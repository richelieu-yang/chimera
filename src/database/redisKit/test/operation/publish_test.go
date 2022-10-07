package operation

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/database/redisKit/test"
	"testing"
)

func TestPublish(t *testing.T) {
	fmt.Println("测试 Publish（发布） -------------------")

	// client
	client, err := test.NewSingleNodeClient()
	if err != nil {
		panic(err)
	}

	err = client.Publish("", "")
	if err != nil {
		panic(err)
	}
	fmt.Println("测试 Publish（发布） -------------------")
}
