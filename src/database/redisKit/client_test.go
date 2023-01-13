package redisKit

import (
	"context"
	"fmt"
	"testing"
)

// TestSingleNodeMode 测试Redis的Cluster集群模式
func TestSingleNodeMode(test *testing.T) {
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

	s, err := client.ScanFully(context.TODO(), "*", 10)
	fmt.Println(s, err)

	//value := true
	//ok, err := client.Set(context.TODO(), "ccc", value, 0)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("ok: [%t].\n", ok)
}

func TestClusterMode(test *testing.T) {
	config := &RedisConfig{
		UserName:           "",
		Password:           "",
		Mode:               ClusterMode,
		SingleNodeConfig:   nil,
		MasterSlaverConfig: nil,
		SentinelConfig:     nil,
		ClusterConfig: &ClusterConfig{
			Addrs: []string{
				"127.0.0.1:6380",
				"127.0.0.1:6381",
				"127.0.0.1:6382",
				"127.0.0.1:6383",
				"127.0.0.1:6384",
				"127.0.0.1:6385",
			},
		},
	}
	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}

	//for i := 0; i < 10; i++ {
	//	_, _ = client.Set(context.TODO(), strconv.Itoa(i), strconv.Itoa(i), 0)
	//}

	for i := 0; i < 100; i++ {
		s, _ := client.ScanFully(context.TODO(), "*", 3)
		//fmt.Println(len(s))
		fmt.Println(s)
		if len(s) != 11 {
			//fmt.Println(666)
			//panic(666)
		}
	}
}
