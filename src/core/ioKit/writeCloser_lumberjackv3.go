package ioKit

import (
	"github.com/natefinch/lumberjack/v3"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"time"
)

type (
	lumberjackOptions struct {
		maxAge     time.Duration
		maxBackups int
		localTime  bool
		compress   bool
	}

	LumberjackOption func(opts *lumberjackOptions)
)

func WithMaxAge(maxAge time.Duration) LumberjackOption {
	return func(opts *lumberjackOptions) {
		opts.maxAge = maxAge
	}
}

func WithMaxBackups(maxBackups int) LumberjackOption {
	return func(opts *lumberjackOptions) {
		opts.maxBackups = maxBackups
	}
}

func WithLocalTime(localTime bool) LumberjackOption {
	return func(opts *lumberjackOptions) {
		opts.localTime = localTime
	}
}

func WithCompress(compress bool) LumberjackOption {
	return func(opts *lumberjackOptions) {
		opts.compress = compress
	}
}

func loadOptions(options ...LumberjackOption) *lumberjackOptions {
	opts := &lumberjackOptions{}
	for _, option := range options {
		option(opts)
	}
	return opts
}

// NewRotatableWriteCloser 可rotate（依据传参maxSize）的io.WriteCloser.
/*
@param maxSize unit: byte
@param options 可选配置:
WithCompress()		[默认不压缩] 是否压缩被rotate的文件？
WithMaxBackups()	[默认保留所有旧日志] 旧日志保存的最多数量
WithMaxAge()		[默认保留所有旧日志] 旧日志保存的最长时间
WithLocalTime()		[默认使用UTC时间] 是否使用本地时间戳？
*/
func NewRotatableWriteCloser(filePath string, maxSize int64, options ...LumberjackOption) (*lumberjack.Roller, error) {
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

	opts := loadOptions(options...)
	return lumberjack.NewRoller(filePath, maxSize, &lumberjack.Options{
		MaxAge:     opts.maxAge,
		MaxBackups: opts.maxBackups,
		LocalTime:  opts.localTime,
		Compress:   opts.compress,
	})
}
