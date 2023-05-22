package main

import (
	"github.com/richelieu42/chimera/v2/src/jsonKit"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := jsonKit.Unmarshal(nil, nil); err != nil {
		logrus.Fatal(err)
	}
}
