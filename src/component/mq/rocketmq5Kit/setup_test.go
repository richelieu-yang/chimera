package rocketmq5Kit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/config/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	type config struct {
		RocketMQ5 *Config `json:"rocketmq5,optional"`
	}

	if wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName); err != nil {
		panic(err)
	} else {
		fmt.Printf("new working directory: [%s].\n", wd)
	}

	c := &config{}
	confKit.MustLoad("_chimera-lib/config.yaml", c)
	fmt.Println(c)
}
