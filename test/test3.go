package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/setKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/sirupsen/logrus"
)

func main() {
	set := setKit.NewSet[string](true)

	go func() {
		logrus.Info(set.)
	}()

	for {
		id := idKit.NewUUID()
		if set.Contains(id) {
			panic("already exists, id: " + id)
		}
		set.Add(id)
	}
}
