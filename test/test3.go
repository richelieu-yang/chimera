package main

import (
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/richelieu-yang/chimera/v2/src/component/nacosKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	ip := "10.0.9.141"
	port := 12345

	nacosKit.MustSetUp(&nacosKit.Config{
		NamespaceId: "",
		Addresses:   []string{"http://localhost:8848/nacos"},
	})
	client, err := nacosKit.NewNamingClient()
	if err != nil {
		logrus.Fatal(err)
	}
	defer client.CloseClient()

	go func() {
		logrus.Info("sleep starts")
		time.Sleep(time.Second * 10)
		logrus.Info("sleep ends")

		client.CloseClient()
		logrus.Info("-")
	}()

	ok, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:       ip,
		Port:     uint64(port),
		Weight:   2000,
		Enable:   true,
		Healthy:  true,
		Metadata: nil,
		//ClusterName: "",
		ServiceName: "wmq",
		//GroupName:   "",
		Ephemeral: false,
	})
	if err != nil {
		panic(err)
	}
	if !ok {
		panic("fail to register")
	}
	for {
	}
}
