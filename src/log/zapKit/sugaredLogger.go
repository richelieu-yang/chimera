package zapKit

import (
	"go.uber.org/zap"
)

func NewSugaredLogger(logger *zap.Logger) *zap.SugaredLogger {
	return logger.Sugar()
}
