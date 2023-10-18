package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	logrusKit.MustSetUp(nil)

	wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}
	logrus.Infof("wd: [%s].", wd)
	path := "_chimera-lib/config.yaml"

	type config struct {
		Nacos *Config `json:"nacos"`
	}

	c := &config{}
	if _, err := viperKit.ReadFileAs(path, nil, c); err != nil {
		panic(err)
	}
	MustSetUp(c.Nacos)

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
