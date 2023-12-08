package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/mq/rocketmq5Kit"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	var (
		//topic = "test"
		tag = "test"
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

	producer, err := rocketmq5Kit.NewProducer()
	if err != nil {
		logrus.Fatal(err)
	}
	ulid := idKit.NewULID()
	for i := 0; ; i++ {
		time.Sleep(time.Second * 3)

		msg := rocketmq5Kit.NewMessage("test", []byte(fmt.Sprintf("%s_%d", ulid, i)), &tag)
		sendReceipts, err := producer.Send(context.TODO(), msg)
		if err != nil {
			logrus.WithError(err).Errorf("Fail to send message(text: %s).", string(msg.Body))
			continue
		}

		logrus.Infof("length: [%d].", len(sendReceipts))
		for _, sendReceipt := range sendReceipts {
			logrus.WithFields(logrus.Fields{
				"MessageID":     sendReceipt.MessageID,
				"TransactionId": sendReceipt.TransactionId,
				"Offset":        sendReceipt.Offset,
			}).Infof("Manager to send message(text: %s).", string(msg.Body))
		}
		logrus.Info("------")
	}
}
