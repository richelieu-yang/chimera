package test

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/core/timeKit"
	"testing"
	"time"
)

func Test(t *testing.T) {
	fmt.Println("测试 ---------------------------------------")

	// client
	client, err := NewSentinelClient()
	if err != nil {
		panic(err)
	}

	err = client.FlushDB()
	fmt.Println(err)

	// code
	cron, err := timeKit.NewCron("@every 10s", func() {
		str := timeKit.FormatCurrentTime(timeKit.DirFormat)
		flag, err := client.Set(str, "ccc", time.Second*15)
		fmt.Println(flag, err)
	})
	if err != nil {
		panic(err)
	}
	cron.Start()

	select {}
}
