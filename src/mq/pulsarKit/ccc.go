package pulsarKit

import (
	"github.com/richelieu42/chimera/src/funcKit"
	"github.com/sirupsen/logrus"
)

func Ccc() {
	logrus.Info(funcKit.GetCaller(1))
}
