package skyWalkingKit

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/netKit"
	"github.com/sirupsen/logrus"
)

func MustSetUp(config *Config) {
	if err := SetUp(config); err != nil {
		logrus.Fatal(err)
	}
}

func SetUp(config *Config) error {
	serverAddr := config.ServerAddr
	addr, err := netKit.ParseToAddress(serverAddr)
	if err != nil {
		return err
	}

	fmt.Println(addr.String())

	return nil
}
