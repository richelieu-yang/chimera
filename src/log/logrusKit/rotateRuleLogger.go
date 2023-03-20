package logrusKit

import (
	"github.com/richelieu42/chimera/src/core/ioKit"
	"github.com/sirupsen/logrus"
)

func NewSizeLimitRotateRuleLogger(filePath, delimiter string, days, maxSize, maxBackups int, compress bool, level logrus.Level) (*logrus.Logger, error) {
	wc, err := ioKit.NewSizeLimitRotateRuleWriteCloser(filePath, delimiter, days, maxSize, maxBackups, compress)
	if err != nil {
		return nil, err
	}

	logger := NewLogger(nil, level)
	logger.Out = wc

	return logger, nil
}

func NewDailyRotateRuleLogger(filePath, delimiter string, days int, compress bool, level logrus.Level) (*logrus.Logger, error) {
	wc, err := ioKit.NewDailyRotateRuleWriteCloser(filePath, delimiter, days, compress)
	if err != nil {
		return nil, err
	}

	logger := NewLogger(nil, level)
	logger.Out = wc

	return logger, nil
}
