package main

import (
	"github.com/richelieu42/chimera/src/mq/pulsarKit"
)

func main() {
	pulsarKit.MustSetUp(&pulsarKit.Config{
		Addresses:      nil,
		LogPath:        "",
		TopicForVerify: "",
	})
}
