package mqKit

import (
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewProducer(t *testing.T) {
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
	c.RocketMQ5.ClientLogPath = "consumer.log"
	MustSetUp(c.RocketMQ5)

	producer, err := NewProducer()
	if err != nil {
		logrus.Fatal(err)
	}
	for i := 0; i < 3; i++ {

	}

}
