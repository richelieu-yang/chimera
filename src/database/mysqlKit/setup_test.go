package mysqlKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/jsonKit"
	"testing"
)

func TestSetUp(t *testing.T) {
	if wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName); err != nil {
		panic(err)
	} else {
		fmt.Printf("wd: %s\n", wd)
	}

	type config struct {
		MySQL *Config `json:"mysql"`
	}

	c := &config{}
	confKit.MustLoad("chimera-lib/config.yaml", c)

	str, err := jsonKit.MarshalToString(c, jsonKit.WithIndent("    "))
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
