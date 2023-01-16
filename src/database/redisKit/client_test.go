package redisKit

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
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

	c := client.GetGoRedisClient()

	sc := c.Scan(context.TODO(), 0, "*", 10)
	iter := sc.Iterator()
	for iter.Next(context.TODO()) {
		fmt.Println(iter.Val())
	}

	//for i := 0; i < 10; i++ {
	//	_, _ = client.Set(context.TODO(), strconv.Itoa(i), strconv.Itoa(i), 0)
	//}
	//for i := 0; i < 100; i++ {
	//	s, _ := client.ScanFully(context.TODO(), "*", 3)
	//	//fmt.Println(len(s))
	//	fmt.Println(s)
	//	if len(s) != 11 {
	//		//fmt.Println(666)
	//		//panic(666)
	//	}
	//}
}

func scan(client redis.UniversalClient, ctx context.Context, match string, count int64) ([]string, error) {
	var cursor uint64 = 0
	var keys []string

	for {
		var s []string
		var err error
		s, cursor, err = client.Scan(ctx, cursor, match, count).Result()
		if err != nil {
			return nil, err
		}

		keys = sliceKit.Merge(keys, s)
		if cursor == 0 {
			// 完整的过一遍了，中断循环
			break
		}
	}
	return sliceKit.RemoveDuplicate(keys), nil
}

func scan1(client redis.UniversalClient, ctx context.Context, match string, count int64) ([]string, error) {
	var cursor uint64 = 0
	var keys []string

	scanCmd := client.Scan(ctx, cursor, match, count)
	iter := scanCmd.Iterator()
	for iter.Next(ctx) {
		tmp := iter.Val()
		keys = append(keys, tmp)
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	if keys == nil {
		keys = []string{}
	} else {
		keys = sliceKit.RemoveDuplicate(keys)
	}
	return keys, nil

	//s, cursor, err = .Result()
	//if err != nil {
	//	return nil, err
	//}
}
