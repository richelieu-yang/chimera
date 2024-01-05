package main

import (
	"flag"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
	"github.com/sirupsen/logrus"
)

var (
	port   *uint
	target *string
)

func init() {
	logrusKit.MustSetUp(&logrusKit.Config{
		Level:      "",
		PrintBasic: true,
	})

	port = flag.Uint("port", 80, "port of service")
	target = flag.String("target", "", "target of proxy")
}

func main() {
	flag.Parse()

	logrus.Infof("port: [%d].", *port)
	logrus.Infof("target: [%s].", *target)

	if err := netKit.AssertValidPort(int(*port)); err != nil {
		panic(err)
	}
	if err := strKit.AssertNotEmpty(*target, "target"); err != nil {
		panic(err)
	}

}
