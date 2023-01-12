package redisKit

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestClusterMode 测试Redis的Cluster集群模式
func TestClusterMode(test *testing.T) {
	config := &RedisConfig{
		UserName: "",
		Password: "",
		Mode:     SingleNodeMode,
		SingleNodeConfig: &SingleNodeConfig{
			Addr: "127.0.0.1:6379",
			DB:   10,
		},
		MasterSlaverConfig: nil,
		SentinelConfig:     nil,
		ClusterConfig:      nil,
	}
	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}

	flag, err := client.Expire(context.TODO(), "a", time.Second*100)
	if err != nil {
		panic(err)
	}
	fmt.Println(flag, err)

	//value := true
	//ok, err := client.Set(context.TODO(), "ccc", value, 0)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("ok: [%t].\n", ok)
}
