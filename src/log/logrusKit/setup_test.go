package logrusKit

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSetUp(t *testing.T) {
	SetUp(&Config{
		Level:      "",
		PrintBasic: true,
	})

	DisableQuoteTemporarily(nil, func() {
		logrus.Info("1\n2\n3\n")
	})
	logrus.Info("1\n2\n3\n")
}
