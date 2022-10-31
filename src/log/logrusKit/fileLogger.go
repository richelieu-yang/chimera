package logrusKit

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// NewFileLogger
func NewFileLogger(logPath string, formatter logrus.Formatter, level logrus.Level, toConsole bool) error {
	var writer io.Writer

	f, err := os.Create(logPath)
	if err != nil {
		return err
	}
	if toConsole {
		writer = io.MultiWriter(f, os.Stdout)
	} else {
		writer = f
	}

}
