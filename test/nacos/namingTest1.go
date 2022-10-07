package main

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/component/nacosKit"
	"gitee.com/richelieu042/go-scales/src/core/timeKit"
	"gitee.com/richelieu042/go-scales/src/randomKit"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

func main() {
	config := &nacosKit.NacosConfig{
		ServerAddresses: []string{"192.168.1.73:8848"},
	}

	client, err := nacosKit.NewNamingClient(config, "d:/nacosTestLog/", nacosKit.DebugLevel)
	if err != nil {
		panic(err)
	}

	// 注册实例
	flag, err := nacosKit.RegisterInstance(client, vo.RegisterInstanceParam{
		Ip:          "2.2.2.2",
		Port:        2222,
		ServiceName: "test",
		Weight:      8000,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   false,
		Metadata:    map[string]string{"idc": "shanghai"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("注册服务结果：", flag)

	timeKit.SetInterval(func(t time.Time) {
		ok, err := nacosKit.UpdateInstance(client, vo.UpdateInstanceParam{
			Ip:          "10.0.0.80",
			Port:        8848,
			ServiceName: "test",
			Weight:      randomKit.Float64(100, 2),
			Enable:      true,
			Ephemeral:   false,
			Metadata:    map[string]string{"idc": "shanghai"},
		})
		fmt.Println(ok, err)
	}, 10*time.Second)

	time.Sleep(time.Minute * 10)
}
