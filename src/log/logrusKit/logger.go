package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/sirupsen/logrus"
	"io"
)

type (
	loggerOptions struct {
		// formatter 日志格式
		formatter logrus.Formatter
		// reportCaller 默认: true
		reportCaller bool
		// level 日志级别，默认: logrus.DebugLevel
		level  logrus.Level
		output io.Writer
		// msgPrefix 日志输出的msg属性的前缀（默认: ""）
		/*
			PS: 如果不为""的话，拼接时会在 msgPrefix 和 msg 间加个空格.
		*/
		msgPrefix string
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

func loadOptions(options ...LoggerOption) *loggerOptions {
	/* 默认值s */
	opts := &loggerOptions{
		formatter:    nil,
		reportCaller: true,
		level:        logrus.DebugLevel,
		// 默认: 输出到控制台
		output:    nil,
		msgPrefix: "",
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
(1) 如果希望 输出到文件 且 rotatable，可以使用 WithOutput()，详见下例.

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

	return logger
}

// NewFileLogger 输出到文件(not rotatable).
func NewFileLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	file, err := fileKit.CreateInAppendMode(filePath)
	if err != nil {
		return nil, err
	}
	option := WithOutput(file)
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
