package cluster

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/database/redisKit/test"
	"github.com/richelieu42/go-scales/src/idKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSingleNode(t *testing.T) {
	fmt.Println("测试单节点集群 -------------------")

	// client
	client, err := test.NewSingleNodeClient()
	if err != nil {
		panic(err)
	}

	// 测试：set、get
	uuid := idKit.NewUUID()
	_, err = client.Set("a", uuid, 0)
	if err != nil {
		panic(err)
	}
	value, err := client.Get("a")
	if err != nil {
		panic(err)
	}
	if value != uuid {
		logrus.Panicf("get的结果(%s)和set的值(%s)不一致！", value, uuid)
	}
}
