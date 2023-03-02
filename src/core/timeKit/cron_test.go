package timeKit

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewCron(t *testing.T) {
	c, _, err := NewCron("@every 1m", func() {
		logrus.Info("---")
	})
	if err != nil {
		logrus.Panic(err)
	}
	c.Start()

	select {}
}
