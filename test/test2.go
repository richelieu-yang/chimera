package main

import (
	"context"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/component/mq/rocketmq5Kit"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	var (
		topic         = "test"
		consumerGroup = "a"
		tag           = "test"
	)

	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			logrus.Fatal(err)
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
		logrus.Fatal(err)
	}
	rocketmq5Kit.MustSetUp(c.RocketMQ5, "_client.log", nil)

	consumer, err := rocketmq5Kit.NewSimpleConsumer(consumerGroup, map[string]*rmq_client.FilterExpression{
		topic: rmq_client.NewFilterExpression(tag),
	})
	if err != nil {
		logrus.Fatal(err)
	}

	for {
		time.Sleep(time.Second)

		mvs, err := consumer.Receive(context.TODO(), rocketmq5Kit.DefaultMaxMessageNum, rocketmq5Kit.DefaultInvisibleDuration)
		if err != nil {
			logrus.WithError(err).Error("Fail to receive message.")
			continue
		}
		for _, mv := range mvs {
			text := string(mv.GetBody())

			err := consumer.Ack(context.TODO(), mv)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"tag":   rocketmq5Kit.GetTagString(mv.GetTag()),
					"text":  text,
					"error": err.Error(),
				}).Error("Fail to ack message.")
				continue
			}
			logrus.WithFields(logrus.Fields{
				"tag":  rocketmq5Kit.GetTagString(mv.GetTag()),
				"text": text,
			}).Info("Manager to receive and ack a message.")
		}
	}
}
