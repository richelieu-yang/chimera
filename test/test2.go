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

			//"172.18.21.32:26380",
			//"172.18.21.32:26381",
			//"172.18.21.32:26382",
			//"172.18.21.32:26383",
			//"172.18.21.32:26384",
			//"172.18.21.32:26385",
		},
		////Password: "123456",
		////DialTimeout:  100 * time.Microsecond,
		//ReadTimeout: 100 * time.Microsecond,
		////WriteTimeout: 100 * time.Microsecond,
		//DialTimeout: 5 * time.Second, //连接建立超时时间，默认5秒。
		////ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		//WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
		//OnConnect: func(ctx context.Context, conn *redis.Conn) error {
		//	fmt.Printf("创建新的连接: %v\n", conn)
		//	return nil
		//},
	})
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	str, err := Cluster.Set(context.Background(), "k", "v", 0).Result()
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
	fmt.Println(str)
}
