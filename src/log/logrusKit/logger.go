package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type (
	loggerOptions struct {
		// formatter 日志格式
		formatter logrus.Formatter

		// reportCaller 默认: true
		reportCaller bool

		// level 日志级别，默认: logrus.DebugLevel
		level logrus.Level

		// output 默认: os.Stderr
		output io.Writer

		// msgPrefix 日志输出的msg属性的前缀，(1) 默认: ""; (2) 非空的话，会拼接在msg前面
		msgPrefix string

		// disableQuote 默认: false
		disableQuote bool
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

func WithMsgPrefix(msgPrefix string) LoggerOption {
	return func(opts *loggerOptions) {
		opts.msgPrefix = msgPrefix
	}
}

func WithDisableQuote(disableQuote bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.disableQuote = disableQuote
	}
}

func loadOptions(options ...LoggerOption) *loggerOptions {
	/* 默认值s */
	opts := &loggerOptions{
		formatter:    nil,
		reportCaller: true,
		level:        logrus.DebugLevel,
		// 默认: 输出到控制台
		output:       nil,
		msgPrefix:    "",
		disableQuote: false,
	}

	for _, option := range options {
		option(opts)
	}

	// 容错，以防"调用方"瞎传
	if opts.formatter == nil {
		opts.formatter = NewDefaultTextFormatter()
	}
	return opts
}

// NewLogger
/*
PS:
(1) 默认输出到 控制台(os.Stderr);
(2) 如果希望 输出到文件 且 rotatable，可以使用 WithOutput()，详见下例.

@param options 可以什么都不配置（此时输出到控制台）
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

	// msgPrefix
	if strKit.IsNotEmpty(opts.msgPrefix) {
		hook := &defaultPrefixHook{prefix: opts.msgPrefix}
		logger.AddHook(hook)
	}

	if opts.disableQuote {
		DisableQuote(logger)
	}

	return logger
}

// NewFileLogger 输出到 文件(not rotatable).
func NewFileLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	file, err := fileKit.CreateInAppendMode(filePath)
	if err != nil {
		return nil, err
	}

	options = append(options, WithOutput(file))
	return NewLogger(options...), nil
}

// NewFileAndStdoutLogger 同时输出到 文件(not rotatable) 和 os.Stdout.
func NewFileAndStdoutLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	f, err := fileKit.CreateInAppendMode(filePath)
	if err != nil {
		return nil, err
	}
	output := ioKit.MultiWriter(f, os.Stdout)

	options = append(options, WithOutput(output))
	return NewLogger(options...), nil
}

// DisposeLogger 释放资源（主要针对文件日志）
func DisposeLogger(logger *logrus.Logger) error {
	if logger == nil {
		return nil
	}

	return ioKit.TryToClose(logger.Out)
}
