package rocketmq5Kit

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/config/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewProducer(t *testing.T) {
	var topic string = "test"
	var tag *string = nil

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
	c.RocketMQ5.ClientLogPath = "producer.log"
	MustSetUp(c.RocketMQ5)

	producer, err := NewProducer()
	if err != nil {
		logrus.Fatal(err)
	}
	//timeStr := timeKit.FormatCurrentTime()
	for i := 0; i < 100000; i++ {
		//text := fmt.Sprintf("%s_%d", timeStr, i)
		text := fmt.Sprintf("%d", i)

		msg := &rmq_client.Message{
			Topic: topic,
			Tag:   tag,
			Body:  []byte(text),
		}
		ctx, _ := context.WithTimeout(context.TODO(), time.Second)
		receipts, err := producer.Send(ctx, msg)
		if err != nil {
			logrus.WithError(err).Error("[PRODUCER] fail to send")
		} else {
			receipt := receipts[0]
			logrus.WithFields(logrus.Fields{
				"text":      text,
				"MessageID": receipt.MessageID,
				//"TransactionId": receipt.TransactionId,
				//"Offset":        receipt.Offset,
			}).Info("[PRODUCER] succeed to send")
		}
		time.Sleep(time.Second)
	}
}
