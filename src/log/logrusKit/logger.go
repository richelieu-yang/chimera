package logrusKit

import (
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/ioKit"
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
/*
@param formatter 可以为nil(此时将采用默认值)
*/
func NewLogger(formatter logrus.Formatter, level logrus.Level) *logrus.Logger {
	if formatter == nil {
		formatter = DefaultTextFormatter
	}

	logger := logrus.New()
	logger.SetFormatter(formatter)
	logger.SetLevel(level)
	return logger
}

// NewFileLogger
/*
PS: 如果 logger.Out 被释放后继续调用 logger 进行输出，会失败（e.g. 控制台os.Stderr有输出: Failed to write to log, write /Users/richelieu/Downloads/a.txt: file already closed）.

@param filePath			(1) 文件不存在，会 尝试创建父级目录 && 创建日志文件；(2) 文件存在，会将内容追加在后面
@param formatter 		可以为nil，此时将采用默认值
@param toConsoleFlag 	true: 输出到日志文件的同时，也输出到控制台; false: 只输出到文件日志
*/
func NewFileLogger(filePath string, formatter logrus.Formatter, level logrus.Level, toConsoleFlag bool) (*logrus.Logger, error) {
	// 尝试创建父级目录
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

	var writeCloser io.WriteCloser
	writeCloser, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return newFileLogger(formatter, level, writeCloser, toConsoleFlag), nil
}

// NewRotateFileLogger
/*
PS: 如果 logger.Out 被释放后继续调用 logger 进行输出，会失败（e.g. 控制台os.Stderr有输出: Failed to write to log, invalid argument）.
*/
func NewRotateFileLogger(filePath string, formatter logrus.Formatter, level logrus.Level, rotationTime, maxAge time.Duration, softLinkFlag, toConsoleFlag bool) (*logrus.Logger, error) {
	writeCloser, err := ioKit.NewRotateFileWriteCloser(filePath, rotationTime, maxAge, softLinkFlag)
	if err != nil {
		return nil, err
	}

	/* 此方法不方便 Close()，因为是通过Hook实现输出到控制台的同时也输出到文件日志 */
	//logger := NewLogger(formatter, level)
	//if toConsoleFlag {
	//	// (1) 输出到: 文件日志 + 控制台
	//	lfsHook := lfshook.NewHook(lfshook.WriterMap{
	//		logrus.TraceLevel: writeCloser,
	//		logrus.DebugLevel: writeCloser,
	//		logrus.InfoLevel:  writeCloser,
	//		logrus.WarnLevel:  writeCloser,
	//		logrus.ErrorLevel: writeCloser,
	//		logrus.FatalLevel: writeCloser,
	//		logrus.PanicLevel: writeCloser,
	//	}, formatter)
	//	logger.AddHook(lfsHook)
	//} else {
	//	// (2) 输出到: 文件日志
	//	logger.Out = writeCloser
	//}
	//return logger, nil

	return newFileLogger(formatter, level, writeCloser, toConsoleFlag), nil
}

// newFileLogger 复用代码
func newFileLogger(formatter logrus.Formatter, level logrus.Level, writeCloser io.WriteCloser, toConsole bool) *logrus.Logger {
	logger := NewLogger(formatter, level)
	if toConsole {
		writeCloser = ioKit.MultiWriteCloser(writeCloser, ioKit.NopWriteCloser(os.Stdout))
	}
	logger.Out = writeCloser
	return logger
}
