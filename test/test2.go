package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

var Cluster *redis.ClusterClient

func main() {
	Cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"127.0.0.1:6380",
			"127.0.0.1:6381",
			"127.0.0.1:6382",
			"127.0.0.1:6383",
			"127.0.0.1:6384",
			"127.0.0.1:6385",
		},
	})
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	str, err := Cluster.Del(context.Background(), "a").Result()
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
	fmt.Println(str)
}
