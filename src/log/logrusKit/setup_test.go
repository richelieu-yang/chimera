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

	DisableQuoteTemporarily(nil, func() {
		logrus.Info("1\n2\n3\n")
	})
	logrus.Info("1\n2\n3\n")
}
