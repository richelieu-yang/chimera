package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/mq/rocketmq5Kit"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/sirupsen/logrus"
)

func main() {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		logrus.Infof("wd: [%s].", wd)
	}

	path := "_chimera-lib/config.yaml"
	type config struct {
		RocketMQ5 *rocketmq5Kit.Config `json:"rocketmq5"`
	}
	c := &config{}
	_, err := viperKit.UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}
	rocketmq5Kit.MustSetUp(c.RocketMQ5, "_client.log", &rocketmq5Kit.VerifyConfig{
		Topic:   "test",
		LogPath: "",
	})

	producer, err := rocketmq5Kit.NewProducer()
	if err != nil {
		panic(err)
	}
	ulid := idKit.NewULID()
	i := 0
	for i := 0; ; i++ {
		text := fmt.Sprintf("%s_%d", ulid, i)
		producer.Send(context.TODO())
	}
}
