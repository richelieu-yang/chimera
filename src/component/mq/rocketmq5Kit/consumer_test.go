package rocketmq5Kit

import (
	"context"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/config/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewSimpleConsumer(t *testing.T) {
	var topic string = "test"

	logrusKit.MustSetUp(nil)
	if wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName); err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Infof("new working directory: [%s].", wd)
	}

	type config struct {
		RocketMQ5 *Config `json:"rocketmq5,optional"`
	}
	c := &config{}
	confKit.MustLoad("_chimera-lib/config.yaml", c)
	c.RocketMQ5.ClientLogPath = "consumer.log"
	MustSetUp(c.RocketMQ5)

	consumer, err := NewSimpleConsumer("cg0222", map[string]*rmq_client.FilterExpression{
		topic: rmq_client.SUB_ALL,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	defer consumer.GracefulStop()
	go func() {
		for {
			time.Sleep(time.Millisecond * 50)

			mvs, err := consumer.Receive(context.TODO(), MaxMessageNum, InvisibleDuration)
			if err != nil {
				logrus.WithError(err).Error("[CONSUMER] fail to receive")
				continue
			}
			for _, mv := range mvs {
				go func(mv *rmq_client.MessageView) {
					if err := consumer.Ack(context.TODO(), mv); err != nil {
						logrus.WithFields(logrus.Fields{
							//"topic": mv.GetTopic(),
							//"tag":   mv.GetTag(),
							//"msgId": mv.GetMessageId(),
							"text":  string(mv.GetBody()),
							"error": err,
						}).Error("[CONSUMER] fail to ack the message")
					}
					logrus.WithFields(logrus.Fields{
						//"topic": mv.GetTopic(),
						//"tag":   mv.GetTag(),
						//"msgId": mv.GetMessageId(),
						"text": string(mv.GetBody()),
					}).Info("[CONSUMER] receive a message")
				}(mv)
			}
		}
	}()
	select {}
}
