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
		Redis *Config `json:"redis"`
	}
	c := &config{}
	path := "chimera-lib/config.yaml"
	confKit.MustLoad(path, c)
	MustSetUp(c.Redis)

	client, err := GetClient()
	if err != nil {
		logrus.Fatal(err)
	}

	//flag, err := client.HExists(context.TODO(), "ccc1", "2")
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//println("HExists:", flag)

	fmt.Println(client.Set(context.TODO(), "222", "222", -1))

}
