package logrusKit

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	MustSetUp(&Config{
		Level:      "debug",
		PrintBasic: true,
	})

	DisableQuote(nil)
	logrus.Info("1\n2\n3\n")

	EnableQuote(nil)
	logrus.Info("1\n2\n3\n")
}
