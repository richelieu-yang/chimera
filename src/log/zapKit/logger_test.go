package zapKit

import (
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewLogger(t *testing.T) {
	// 日志文件会生成在当前.go文件的同目录下
	writer, err := ioKit.NewLumberjackWriteCloser(ioKit.WithFilePath("aaa.log"))
	assert.Nil(t, err)

	logger, err := NewLogger(writer, zap.InfoLevel)
	assert.Nil(t, err)

	// 由于日志级别，下一行不会输出
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
