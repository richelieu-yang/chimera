package zapKit

import (
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger
/*
@param level e.g. zap.InfoLevel
*/
func NewLogger(logPath string, maxFileSize, maxFileIndex int, compress bool, level zapcore.Level) (*zap.Logger, error) {
	if err := fileKit.MkParentDirs(logPath); err != nil {
		return nil, err
	}

	encoder := getEncoder()

	writeSyncer := getLogWriter(logPath, maxFileSize, maxFileIndex, compress)

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(encoder, writeSyncer, atomicLevel)
	return zap.New(core, zap.AddCaller()), nil
}

func getLogWriter(logPath string, maxFileSize, maxFileIndex int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    maxFileSize,
		MaxBackups: maxFileIndex,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
