package timeKit

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestSetInterval(t *testing.T) {
	i := SetInterval(func(t time.Time) {
		logrus.Info(t)
	}, time.Second*3)

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")

	ClearInterval(i)
	ClearInterval(i)
	ClearInterval(i)

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")
}
