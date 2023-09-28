package nacosKit

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	config := &Config{
		NamespaceId: "6393fde0-464c-43e1-9c4e-60f4f372a74f",
		Addresses:   []string{"http://localhost:8848/nacos"},
	}

	MustSetUp(config)

	namingClient, err := NewNamingClient()
	if err != nil {
		panic(err)
	}
	defer namingClient.CloseClient()
	configClient, err := NewConfigClient()
	if err != nil {
		panic(err)
	}
	defer configClient.CloseClient()

	ok, err := namingClient.UpdateInstance(vo.UpdateInstanceParam{
		Ip:   "127.0.0.1",
		Port: 80,
		/*
			有效范围: [0.0, 10000, 0]
		*/
		Weight:      10000.01,
		Enable:      true,
		Healthy:     true,
		Metadata:    nil,
		ClusterName: "",
		ServiceName: "ws",
		GroupName:   "wo3",
		Ephemeral:   false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("ok", ok)
}
