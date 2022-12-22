package pipeline

import (
	"context"
	"fmt"
	"github.com/richelieu42/go-scales/src/database/redisKit/test"
	"testing"
)

func TestPipeline(t *testing.T) {
	fmt.Println("测试 Pipeline（管道） -------------------")

	// client
	client, err := test.NewSentinelClient()
	if err != nil {
		panic(err)
	}

	pipe := client.Pipeline()
	cmd := pipe.Set(context.TODO(), "a", "c", 0)
	// 此处会报错
	cmd1 := pipe.Do(context.TODO(), "ino")
	cmd2 := pipe.Set(context.TODO(), "b", "c", 0)
	cmd3 := pipe.Do(context.TODO(), "ino111")

	_, err = pipe.Exec(context.TODO())
	fmt.Println(err)

	fmt.Println(cmd.Result())
	fmt.Println(cmd1.Result())
	fmt.Println(cmd2.Result())
	fmt.Println(cmd3.Result())
}
