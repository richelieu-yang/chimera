package logrusKit

import (
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/sirupsen/logrus"
)

// NewLumberjackLogger
/*
@param logger NewBasicLogger
*/
func NewLumberjackLogger(logger *logrus.Logger, filePath string, options ...ioKit.LumberjackOption) (*logrus.Logger, error) {
	writeCloser, err := ioKit.NewLumberjackWriteCloser(filePath, options...)
	if err != nil {
		return nil, err
	}
	logger.SetOutput(writeCloser)
	return logger, nil
}
