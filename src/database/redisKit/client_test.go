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

	if _, err := client.Ping(context.TODO()); err != nil {
		panic(err)
	}

	//for i := 0; i < 1; i++ {
	//	s, err := scan(client.GetGoRedisClient(), context.TODO(), "*", 10)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(s)
	//	fmt.Println(len(s))
	//	fmt.Println("======")
	//
	//	if len(s) != 101 {
	//		panic(666)
	//	}
	//}
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
				//"127.0.0.1:6380",
				//"127.0.0.1:6381",
				//"127.0.0.1:6382",
				//"127.0.0.1:6383",
				//"127.0.0.1:6384",
				//"127.0.0.1:6385",

				"192.168.80.43:7000",
				"192.168.80.43:7001",
				"192.168.80.27:7002",
				"192.168.80.27:7003",
				"192.168.80.42:7004",
				"192.168.80.42:7005",
			},
		},
	}
	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}

	if str, err := client.Ping(context.TODO()); err != nil {
		panic(err)
	} else {
		fmt.Println(str)
	}

	//for i := 0; i < 100; i++ {
	//	_, _ = client.Set(context.TODO(), strconv.Itoa(i), strconv.Itoa(i), 0)
	//}

	//for i := 0; i < 1000; i++ {
	//	s, err := client.ScanFully(context.TODO(), "*", 10)
	//	if err != nil {
	//		panic(err)
	//	}
	//	if len(s) != 101 {
	//		panic(len(s))
	//	}
	//	fmt.Printf("====== %d\n", len(s))
	//}

	//c := client.GetGoRedisClient()

	//sc := c.Scan(context.TODO(), 0, "*", 10)
	//iter := sc.Iterator()
	//for iter.Next(context.TODO()) {
	//	fmt.Println(iter.Val())
	//}
}
