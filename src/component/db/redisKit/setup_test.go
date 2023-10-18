package redisKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"testing"
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
	if _, err := viperKit.ReadFileAs(path, nil, c); err != nil {
		panic(err)
	}

	MustSetUp(c.Redis)
	client, err := GetClient()
	if err != nil {
		logrus.Fatal(err)
	}
	client = client

	{
		fmt.Println(client.XDel(context.TODO(), "tickets", "1697005411917-0"))
		fmt.Println(client.XDel(context.TODO(), "tickets", "1697005411917-0"))
	}
}
