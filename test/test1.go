package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "logs/a/rocketmq_client_go.log",
		MaxSize:    maxFileSize,
		MaxBackups: maxFileIndex,
		Compress:   false,
	}

	//maxFileIndex := 10
	//maxFileSize := 1073741824
	//lumberJackLogger := &lumberjack.Logger{
	//	Filename:   "logs/a/rocketmq_client_go.log",
	//	MaxSize:    maxFileSize,
	//	MaxBackups: maxFileIndex,
	//	Compress:   false,
	//}
	//writeSyncer := zapcore.AddSync(lumberJackLogger)
	//
	//encoder := getEncoder()
	//
	//var atomicLevel = zap.NewAtomicLevel()
	//atomicLevel.SetLevel(zap.InfoLevel)
	//
	//core := zapcore.NewCore(encoder, writeSyncer, atomicLevel)
	//logger := zap.New(core, zap.AddCaller())
	//sugarBaseLogger := logger.Sugar()
	//
	//sugarBaseLogger.Infof("[TEST] %d", 666)
}
