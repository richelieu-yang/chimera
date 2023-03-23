package logrusKit

import (
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// DisposeLogger 释放资源（主要针对文件日志）
func DisposeLogger(logger *logrus.Logger) error {
	if logger == nil {
		return nil
	}
	return ioKit.CloseWriters(logger.Out)
}

// NewLogger 输出到控制台（os.Stderr） + TextFormatter
/*
@param formatter 可以为nil(此时将采用默认值)
*/
func NewLogger(formatter logrus.Formatter, level logrus.Level) *logrus.Logger {
	logger := logrus.New()

	if formatter == nil {
		formatter = NewTextFormatter("")
	}
	logger.SetFormatter(formatter)
	logger.SetReportCaller(true)

	logger.SetLevel(level)
	return logger
}

// NewFileLogger
/*
PS: 如果 logger.Out 被释放后继续调用 logger 进行输出，会失败（e.g. 控制台os.Stderr有输出: Failed to write to log, write /Users/richelieu/Downloads/a.txt: file already closed）.

@param filePath			内部会做处理:
						(1) 会尝试创建父级目录
						(2) 文件不存在，会自动创建
						(3) 文件存在：是个文件，追加在最后；是个目录，返回error
@param formatter 		可以为nil，此时将采用默认值
@param toConsoleFlag 	true: 输出到日志文件的同时，也输出到控制台; false: 只输出到文件日志
*/
func NewFileLogger(filePath string, formatter logrus.Formatter, level logrus.Level, toConsoleFlag bool) (*logrus.Logger, error) {
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}

	writeCloser, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return newFileLogger(formatter, level, writeCloser, toConsoleFlag), nil
}

// newFileLogger 复用代码
func newFileLogger(formatter logrus.Formatter, level logrus.Level, writeCloser io.WriteCloser, toConsole bool) *logrus.Logger {
	logger := NewLogger(formatter, level)
	if toConsole {
		writeCloser = ioKit.MultiWriteCloser(writeCloser, ioKit.NopWriteCloser(os.Stdout))
	}
	logger.SetOutput(writeCloser)
	return logger
}
