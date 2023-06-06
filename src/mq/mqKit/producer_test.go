package mqKit

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewProducer(t *testing.T) {
	var topic string = "test"
	var tag *string = nil

	type config struct {
		RocketMQ5 *Config `json:"rocketmq5,optional"`
	}

	logrusKit.MustSetUp(nil)
	if wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName); err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Infof("new working directory: [%s].\n", wd)
	}

	c := &config{}
	confKit.MustLoad("chimera-lib/config.yaml", c)
	c.RocketMQ5.ClientLogPath = "producer.log"
	MustSetUp(c.RocketMQ5)

	producer, err := NewProducer()
	if err != nil {
		logrus.Fatal(err)
	}
	time := timeKit.FormatCurrentTime()
	for i := 0; i < 3; i++ {
		text := fmt.Sprintf("%s_%d", time, i)

		msg := &rmq_client.Message{
			Topic: topic,
			Tag:   tag,
			Body:  []byte(text),
		}
		receipts, err := producer.Send(context.Background(), msg)
		if err != nil {
			logrus.WithError(err).Error("[PRODUCER] fail to send")
		} else {
			receipt := receipts[0]
			logrus.WithFields(logrus.Fields{
				"MessageID":     receipt.MessageID,
				"TransactionId": receipt.TransactionId,
				"Offset":        receipt.Offset,
			}).Info("[PRODUCER] succeed to send")
		}
	}
}
