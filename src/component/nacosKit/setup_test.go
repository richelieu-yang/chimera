package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	config := &Config{
		NamespaceId: "",
		Addresses:   []string{"http://localhost:8848/nacos"},
	}

	MustSetUp(config)

	client, err := NewNamingClient()
	if err != nil {
		panic(err)
	}
	defer client.CloseClient()
	ok, err := client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:       "127.0.0.1",
		Port:     12345,
		Weight:   100,
		Enable:   true,
		Healthy:  true,
		Metadata: nil,
		//ClusterName: "",
		ServiceName: "test-service",
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

	//client.UpdateInstance(vo.UpdateInstanceParam{
	//	Ip:          "",
	//	Port:        0,
	//	Weight:      0,
	//	Enable:      false,
	//	Healthy:     false,
	//	Metadata:    nil,
	//	ClusterName: "",
	//	ServiceName: "",
	//	GroupName:   "",
	//	Ephemeral:   false,
	//})
}
