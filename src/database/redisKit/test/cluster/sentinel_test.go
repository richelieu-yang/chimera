package cluster

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/database/redisKit/test"
	"gitee.com/richelieu042/go-scales/src/idKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSentinel(t *testing.T) {
	fmt.Println("测试哨兵集群 -------------------")

	// client
	client, err := test.NewSentinelClient()
	if err != nil {
		panic(err)
	}

	// 测试：set、get
	uuid := idKit.NewUUID()
	_, err = client.Set("1", uuid, 0)
	if err != nil {
		panic(err)
	}
	value, err := client.Get("1")
	if err != nil {
		panic(err)
	}
	if value != uuid {
		logrus.Panicf("get的结果(%s)和set的值(%s)不一致！", value, uuid)
	}
}
