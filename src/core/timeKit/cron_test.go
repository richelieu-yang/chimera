package timeKit

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestNewCron(t *testing.T) {
	logrus.Info(time.Now())
	c, _, err := NewCron("@every 10s", func() {
		logrus.Info(time.Now())
	})
	if err != nil {
		logrus.Panic(err)
	}
	c.Start()

	select {}
}
