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

	for i := 0; i < 1; i++ {
		s, err := scan(client.GetGoRedisClient(), context.TODO(), "*", 10)
		if err != nil {
			panic(err)
		}
		fmt.Println(s)
		fmt.Println(len(s))
		fmt.Println("======")

		if len(s) != 101 {
			panic(666)
		}
	}
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

	//for i := 0; i < 100; i++ {
	//	_, _ = client.Set(context.TODO(), strconv.Itoa(i), strconv.Itoa(i), 0)
	//}

	for i := 0; i < 100; i++ {
		s, err := client.ScanFully(context.TODO(), "*", 10)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(s))
		if len(s) != 101 {
			panic("cccccccccccccccccccc")
		}
	}

	//c := client.GetGoRedisClient()

	//sc := c.Scan(context.TODO(), 0, "*", 10)
	//iter := sc.Iterator()
	//for iter.Next(context.TODO()) {
	//	fmt.Println(iter.Val())
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
		keys = append(keys, iter.Val())
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
}
