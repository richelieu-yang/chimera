package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type (
	loggerOptions struct {
		formatter    logrus.Formatter
		reportCaller bool
		level        logrus.Level
		writer       io.Writer
	}

	LoggerOption func(opts *loggerOptions)
)

func WithFormatter(formatter logrus.Formatter) LoggerOption {
	return func(opts *loggerOptions) {
		opts.formatter = formatter
	}
}

func WithReportCaller(reportCaller bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.reportCaller = reportCaller
	}
}

func WithLevel(level logrus.Level) LoggerOption {
	return func(opts *loggerOptions) {
		opts.level = level
	}
}

func WithWriter(writer io.Writer) LoggerOption {
	return func(opts *loggerOptions) {
		opts.writer = writer
	}
}

func loadOptions(options ...LoggerOption) *loggerOptions {
	/* 默认值s */
	opts := &loggerOptions{
		formatter:    nil,
		reportCaller: true,
		// 默认: debug
		level: logrus.DebugLevel,
		// 默认: 输出到控制台
		writer: nil,
	}

	for _, option := range options {
		option(opts)
	}

	// 容错，以防 调用方 瞎传
	if opts.formatter == nil {
		opts.formatter = DefaultTextFormatter
	}

	return opts
}

// NewLogger
/*
PS:
(1) 如果希望 输出到文件 且 rotatable，可以使用 WithWriter()，详见下例.

@param options 可以什么都不配置（此时输出到控制台）

e.g.
	writer, err := ioKit.NewLumberjackWriteCloser(path)
	// process err
	logger := NewLogger(WithWriter(writer))
*/
func NewLogger(options ...LoggerOption) *logrus.Logger {
	opts := loadOptions(options...)

	logger := logrus.New()
	logger.SetFormatter(opts.formatter)
	logger.SetReportCaller(opts.reportCaller)
	logger.SetLevel(opts.level)
	if opts.writer != nil {
		logger.SetOutput(opts.writer)
	}
	return logger
}

// NewFileLogger 输出到文件(not rotatable).
func NewFileLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	option := WithWriter(file)
	options = append(options, option)
	return NewLogger(options...), nil
}

// DisposeLogger 释放资源（主要针对文件日志）
func DisposeLogger(logger *logrus.Logger) error {
	if logger == nil {
		return nil
	}
	return ioKit.TryToClose(logger.Out)
}
