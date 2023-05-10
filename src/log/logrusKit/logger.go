package logrusKit

import (
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"io"
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
	opts := &loggerOptions{
		formatter:    nil,
		reportCaller: true,
		level:        logrus.DebugLevel,
		writer:       nil,
	}
	for _, option := range options {
		option(opts)
	}

	/* 容错，放在下面以防调用方瞎搞 */
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
	writer, err := ioKit.NewLumberjackWriteCloser(ioKit.WithFilePath(path))
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

// DisposeLogger 释放资源（主要针对文件日志）
func DisposeLogger(logger *logrus.Logger) error {
	if logger == nil {
		return nil
	}
	return ioKit.TryToClose(logger.Out)
}
