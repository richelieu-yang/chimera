package timeKit

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestSetInterval(t *testing.T) {
	ticker := SetInterval(func(t time.Time) {
		logrus.Info(t.String())
	}, time.Second*3)

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")

	ClearInterval(ticker)
	ClearInterval(ticker)
	ClearInterval(ticker)

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")
}
