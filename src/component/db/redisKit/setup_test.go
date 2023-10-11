package redisKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
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
	println("wd:", wd)

	type config struct {
		Redis Config `json:"redis"`
	}

	c := &config{}
	path := "_chimera-lib/config.yaml"
	confKit.MustLoad(path, c)
	MustSetUp(&c.Redis)
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
