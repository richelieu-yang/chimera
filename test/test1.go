package main

import (
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/richelieu42/chimera/v2/src/database/redisKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	path := "/Users/richelieu/GolandProjects/chimera/chimera-lib/config.yaml"

	logrusKit.MustSetUp(nil)

	type config struct {
		Redis *redisKit.Config `json:"redis"`
	}
	c := &config{}
	confKit.MustLoad(path, c)
	redisKit.MustSetUp(c.Redis)

	client, err := redisKit.GetClient()
	if err != nil {
		logrus.Fatal(err)
	}

	//go func() {
	//
	//}()

	mu := client.NewDistributedMutex("/ccc")
	logrus.Info("ready to lock on")
	if err := mu.Lock(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("lock on")

	time.Sleep(time.Second * 9)

	ok, err := mu.Unlock()
	logrus.WithFields(logrus.Fields{
		"ok":    ok,
		"error": err,
	}).Info("lock off")
}
