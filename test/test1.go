package main

import (
	"flag"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/mq/pulsarKit"
	"github.com/sirupsen/logrus"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", "127.0.0.1:6650", "地址")
}

func main() {
	flag.Parse()

	logrusKit.MustSetUp(&logrusKit.Config{
		PrintBasic: false,
	})
	logrus.Infof("addr: [%s]", addr)

	pulsarConfig := &pulsarKit.Config{
		Addresses: []string{
			addr,
		},
		VerifyConfig: &pulsarKit.VerifyConfig{
			Topic: "test",
			Print: true,
		},
	}
	pulsarKit.MustSetUp(pulsarConfig)
}
