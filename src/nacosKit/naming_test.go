package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/richelieu-yang/chimera/v2/src/json/jsoniterKit"
	"github.com/richelieu-yang/chimera/v2/src/randomKit"
	"log"
	"testing"
	"time"
)

var nacosConfig = &NacosConfig{
	ServerAddresses: []string{"192.168.0.159:8848"},
	NamespaceId:     "",
}

func Test(t *testing.T) {
	type arg struct {
		name string
	}
	args := []arg{
		{
			name: "0",
		},
	}

	for _, arg := range args {
		t.Run(arg.name, func(t *testing.T) {
			client, err := NewNamingClient(nacosConfig, "d:/output", DebugLevel)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			param := vo.RegisterInstanceParam{
				Ip:          "127.0.0.1",
				Port:        6666,
				Weight:      0,
				Enable:      randomKit.Bool(),
				Healthy:     true,
				Metadata:    nil,
				ClusterName: "",
				ServiceName: "testServiceName",
				GroupName:   "",
				Ephemeral:   false,
			}

			err = client.Subscribe(&vo.SubscribeParam{
				ServiceName: param.ServiceName,
				Clusters:    nil,
				GroupName:   param.GroupName,
				SubscribeCallback: func(services []model.SubscribeService, err error) {
					if err != nil {
						t.Errorf("监听到错误: %v", err)
					} else {
						t.Logf("----------------------- %v", services)
					}
				},
			})
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			{
				param.Enable = true
				json, _ := jsoniterKit.MarshalToString(param)
				t.Logf("即将注册服务: %s", json)
				ok, err := client.RegisterInstance(param)
				if err != nil {
					t.Errorf("注册服务失败，error: %v", err)
					t.FailNow()
				}
				t.Logf("注册服务的结果: [%t].", ok)
			}
			time.Sleep(time.Second * 5)
			{
				param.Enable = false
				json, _ := jsoniterKit.MarshalToString(param)
				t.Logf("即将注册服务: %s", json)
				ok, err := client.RegisterInstance(param)
				if err != nil {
					t.Errorf("注册服务失败，error: %v", err)
					t.FailNow()
				}
				t.Logf("注册服务的结果: [%t].", ok)
			}

			instances, err := client.SelectInstances(vo.SelectInstancesParam{
				Clusters:    nil,
				ServiceName: "testServiceName",
				GroupName:   "",
				HealthyOnly: false,
			})
			log.Println(instances, err)
		})
	}
}
