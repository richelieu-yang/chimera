package statKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func MustSetup(logPath string) {
	if err := Setup(logPath); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func Setup(logPath string) error {
	if err := fileKit.AssertNotExistOrIsFile(logPath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(logPath); err != nil {
		return err
	}
	f, err := fileKit.NewFileInAppendMode(logPath)
	if err != nil {
		return err
	}
	logger = logrusKit.NewLogger(logrusKit.WithOutput(f))
	logrusKit.DisableQuote(logger)

	c, _, err := cronKit.NewCronWithTask("@every 15s", func() {
		PrintStats(logger)
	})
	c.Start()

	return nil
}
