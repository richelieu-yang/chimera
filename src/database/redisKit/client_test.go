package redisKit

import (
	"context"
	"fmt"
	"testing"
)

// TestClusterMode 测试Redis的Cluster集群模式
func TestClusterMode(test *testing.T) {
	config := &RedisConfig{
		UserName: "",
		Password: "",
		Mode:     SingleNodeMode,
		SingleNodeConfig: &SingleNodeConfig{
			Addr: "",
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

	tmp := false
	if ok, err := client.Set(context.TODO(), "ccc", tmp, 0); err != nil {
		panic(err)
	} else {
		fmt.Printf("ok: [%t].\n", ok)
	}
}
