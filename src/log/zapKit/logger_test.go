package zapKit

import (
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewLogger(t *testing.T) {
	writer, err := ioKit.NewLumberjackWriteCloser(ioKit.WithFilePath("aaa.log"))
	assert.Nil(t, err)

	logger, err := NewLogger(writer, zap.DebugLevel)
	assert.Nil(t, err)

	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
