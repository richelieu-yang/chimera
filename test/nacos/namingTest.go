package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/richelieu42/go-scales/src/component/nacosKit"
	"github.com/richelieu42/go-scales/src/core/floatKit"
)

func main() {
	config := &nacosKit.NacosConfig{
		ServerAddresses: []string{"192.168.1.234:8848"},
	}

	client, err := nacosKit.NewNamingClient(config, "d:/nacosTestLog/", nacosKit.DebugLevel)
	if err != nil {
		panic(err)
	}

	// 注册实例
	flag, err := nacosKit.RegisterInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.80",
		Port:        8848,
		ServiceName: "test",
		Weight:      8000,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // default value is DEFAULT
		GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("注册服务结果：", flag)

	// 注册实例1
	flag1, err := nacosKit.RegisterInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.30",
		Port:        8849,
		ServiceName: "test",
		Weight:      1000,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // default value is DEFAULT
		GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("注册服务结果1：", flag1)

	fmt.Println("------------------------------------------------------------------------------------------------")

	count := 0
	for i := 0; i < 1000; i++ {
		instance, err := nacosKit.SelectOneHealthyInstance(client, vo.SelectOneHealthInstanceParam{
			ServiceName: "test",
			GroupName:   "group-a",             // 默认值DEFAULT_GROUP
			Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
		})
		if err != nil {
			fmt.Println(err.Error())
		} else {
			if instance.Weight == 1000 {
				count++
			}
			fmt.Println(instance)
		}
	}

	fmt.Println("比例：", floatKit.Div(float64(count), 1000))
}
