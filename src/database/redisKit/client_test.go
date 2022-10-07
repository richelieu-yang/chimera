package redisKit

import (
	"gitee.com/richelieu042/go-scales/src/core/timeKit"
	"gitee.com/richelieu042/go-scales/src/idKit"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

// TestClusterMode 测试Redis的Cluster集群模式
func TestClusterMode(test *testing.T) {
	//err := microKit.InitializeRedisComponent("env/micro.yaml")
	//if err != nil {
	//	panic(err)
	//}

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

	//ticker := timeKit.SetInterval(func(t time.Time) {
	//	_, err = client.Set(idKit.NewUUID(), timeKit.FormatCurrentTimeByDefault(), time.Second*10)
	//	if err != nil {
	//		panic(err)
	//	}
	//}, time.Second*3)
	//time.Sleep(time.Second * 60)
	//timeKit.ClearInterval(ticker)

	_ = timeKit.SetInterval(func(t time.Time) {
		key := idKit.NewUUID()
		value := timeKit.FormatCurrentTimeByDefault()

		_, err = client.Set(key, value, time.Second*10)
		if err != nil {
			logrus.Error(err.Error())
		} else {
			logrus.Infof("Item(key: %s, value: %s) is inserted.", key, value)
		}
	}, time.Second)
	select {}
}
