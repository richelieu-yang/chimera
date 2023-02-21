package logrusKit

import (
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/ioKit"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

// DisposeLogger 释放资源（主要针对文件日志）
func DisposeLogger(logger *logrus.Logger) error {
	if logger == nil {
		return nil
	}
	return ioKit.CloseWriters(logger.Out)
}

// NewLogger 输出到控制台（os.Stderr）
func NewLogger() *logrus.Logger {
	return NewCustomizedLogger(nil, logrus.DebugLevel)
}

// NewCustomizedLogger 输出到控制台（os.Stderr）
/*
@param formatter 可以为nil，此时将采用默认值
*/
func NewCustomizedLogger(formatter logrus.Formatter, level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	if formatter == nil {
		formatter = DefaultTextFormatter
	}
	logger.SetFormatter(formatter)
	logger.SetLevel(level)
	return logger
}

// NewFileLogger
/*
PS: 如果 logger.Out 被释放后继续调用 logger 进行输出，会失败（e.g. 控制台os.Stderr有输出: Failed to write to log, write /Users/richelieu/Downloads/a.txt: file already closed）.

@param logPath			(1) 文件不存在，会 尝试创建父级目录 && 创建日志文件；(2) 文件存在，会将内容追加在后面
@param formatter 		可以为nil，此时将采用默认值
@param toConsoleFlag 	true: 输出到日志文件的同时，也输出到控制台; false: 只输出到文件日志
*/
func NewFileLogger(filePath string, formatter logrus.Formatter, level logrus.Level, toConsoleFlag bool) (*logrus.Logger, error) {
	// 尝试创建父级目录
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

	var out io.WriteCloser
	out, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	if toConsoleFlag {
		out = ioKit.MultiWriteCloser(out, ioKit.NopWriteCloser(os.Stdout))
	}

	logger := NewCustomizedLogger(formatter, level)
	logger.Out = out
	return logger, nil
}

// NewRotateFileLogger
/*
PS: 如果 logger.Out 被释放后继续调用 logger 进行输出，会失败（e.g. 控制台os.Stderr有输出: Failed to write to log, invalid argument）.
*/
func NewRotateFileLogger(filePath string, formatter logrus.Formatter, level logrus.Level, rotationTime, maxAge time.Duration, toConsoleFlag bool) (*logrus.Logger, error) {
	writeCloser, err := ioKit.NewRotateFileWriteCloser(filePath, rotationTime, maxAge, true)
	if err != nil {
		return nil, err
	}

	logger := NewCustomizedLogger(formatter, level)
	if toConsoleFlag {
		// (1) 输出到: 文件日志 + 控制台
		lfsHook := lfshook.NewHook(lfshook.WriterMap{
			logrus.TraceLevel: writeCloser,
			logrus.DebugLevel: writeCloser,
			logrus.InfoLevel:  writeCloser,
			logrus.WarnLevel:  writeCloser,
			logrus.ErrorLevel: writeCloser,
			logrus.FatalLevel: writeCloser,
			logrus.PanicLevel: writeCloser,
		}, formatter)
		logger.AddHook(lfsHook)
	} else {
		// (2) 输出到: 文件日志
		logger.Out = writeCloser
	}
	return logger, nil
}
