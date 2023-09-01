package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/setKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	set := setKit.NewSet[string](true)

	go func() {
		for {
			logrus.Info(set.Cardinality())
			time.Sleep(time.Minute)
		}
	}()

	for i := 0; i < 100; i++ {
		go func() {
			for {
				id := idKit.NewUUID()
				if set.Contains(id) {
					panic("already exists, id: " + id)
				}
				set.Add(id)
			}
		}()
	}

	for {
	}
}
