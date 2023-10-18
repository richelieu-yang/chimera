package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	logrusKit.MustSetUp(nil)

	config := &Config{
		NamespaceId: "6393fde0-464c-43e1-9c4e-60f4f372a74f",
		Addrs: []string{
			"http://localhost:8848/nacos",
		},
	}

	MustSetUp(config)

	namingClient, err := NewNamingClient()
	if err != nil {
		logrus.Fatal(err)
	}
	defer namingClient.CloseClient()
	configClient, err := NewConfigClient()
	if err != nil {
		logrus.Fatal(err)
	}
	defer configClient.CloseClient()

	flag, err := namingClient.UpdateInstance(vo.UpdateInstanceParam{
		Ip:   "127.0.0.1",
		Port: 80,
		/*
			有效范围: [0.0, 10000, 0]
		*/
		Weight:      80,
		Enable:      true,
		Healthy:     true,
		Metadata:    nil,
		ClusterName: "",
		ServiceName: "ws",
		GroupName:   "wo3",
		Ephemeral:   false,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("flag: [%t].", flag)
}
