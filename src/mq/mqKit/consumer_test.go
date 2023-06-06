package mqKit

import (
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"testing"
)

func TestNewSimpleConsumer(t *testing.T) {
	type config struct {
		RocketMQ5 *Config `json:"rocketmq5,optional"`
	}

	if wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName); err != nil {
		panic(err)
	} else {
		fmt.Printf("new working directory: [%s].\n", wd)
	}

	c := &config{}
	confKit.MustLoad("chimera-lib/config.yaml", c)

	NewSimpleConsumer("cg01", map[string]*rmq_client.FilterExpression{})

}
