package main

import (
	"github.com/richelieu-yang/chimera/v2/src/mq/pulsarKit"
)

func main() {
	pulsarConfig := &pulsarKit.Config{
		Addresses: []string{
			"192.168.1.128:6650",
		},
		VerifyConfig: &pulsarKit.VerifyConfig{
			Topic: "test",
			Print: true,
		},
	}
	pulsarKit.MustSetUp(pulsarConfig)
}
