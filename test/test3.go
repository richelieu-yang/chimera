package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/database/redisKit"
)

type SampleStruct struct {
	Value string
}

func main() {
	redisKit.MustSetUp(&redisKit.Config{
		MinIdleConns: 64,
		MaxIdleConns: 256,
		PoolSize:     512,
		Mode:         redisKit.ModeCluster,
		ClusterConfig: &redisKit.ClusterConfig{
			Addrs: []string{
				"192.168.80.43:7000",
				"192.168.80.43:7001",
				"192.168.80.27:7002",
				"192.168.80.27:7003",
				"192.168.80.42:7004",
				"192.168.80.42:7005",
			},
		},
	})

	client, err := redisKit.GetClient()
	if err != nil {
		panic(err)
	}
	keys, err := client.ScanFully(context.TODO(), "*", 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
}
