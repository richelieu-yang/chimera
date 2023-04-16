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
		output       io.Writer
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

func WithOutput(output io.Writer) LoggerOption {
	return func(opts *loggerOptions) {
		opts.output = output
	}
}

func loadOptions(options ...LoggerOption) *loggerOptions {
	opts := &loggerOptions{
		formatter:    nil,
		reportCaller: true,
		level:        logrus.DebugLevel,
		output:       nil,
	}
	for _, option := range options {
		option(opts)
	}

	/* 容错，以防调用方瞎搞 */
	if opts.formatter == nil {
		opts.formatter = DefaultTextFormatter
	}

	return opts
}

// NewLogger
/*
@param options 可以什么都不配置
*/
func NewLogger(options ...LoggerOption) *logrus.Logger {
	opts := loadOptions(options...)

	logger := logrus.New()
	logger.SetFormatter(opts.formatter)
	logger.SetReportCaller(opts.reportCaller)
	logger.SetLevel(opts.level)
	if opts.output != nil {
		logger.SetOutput(opts.output)
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
