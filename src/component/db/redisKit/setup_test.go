package redisKit

import (
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestSetUp(t *testing.T) {
	wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}
	logrus.Infof("wd: [%s].", wd)
	path := "_chimera-lib/config.yaml"

	type config struct {
		Redis *Config `json:"redis"`
	}

	c := &config{}
	if _, err := viperKit.UnmarshalFromFile(path, nil, c); err != nil {
		panic(err)
	}

	MustSetUp(c.Redis)
	client, err := GetClient()
	if err != nil {
		logrus.Fatal(err)
	}
	client = client

	{
		mu := client.NewDistributedMutex("aaa")
		err := mu.Lock()
		if err != nil {
			panic(err)
		}
		defer mu.Unlock()

		time.Sleep(time.Second * 3)
	}
}
