package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewSugaredLogger(logPath string, maxFileSize, maxFileIndex int, compress bool, level zapcore.Level) (*zap.SugaredLogger, error) {
	logger, err := NewLogger(logPath, maxFileSize, maxFileIndex, compress, level)
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}
