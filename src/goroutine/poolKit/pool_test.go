package poolKit

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewPool(t *testing.T) {
	//size := 1
	size := 3

	pool, err := NewPool(size)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		j := i
		err := pool.Submit(func() {
			time.Sleep(time.Second)
			logrus.Info(j)
		})
		if err != nil {
			logrus.WithError(err).Error("fail to submit")
		}
	}
	for {
		logrus.WithFields(logrus.Fields{
			"Running": pool.Running(),
			"Waiting": pool.Waiting(),
		}).Info("-")

		time.Sleep(time.Millisecond * 100)
	}
}
