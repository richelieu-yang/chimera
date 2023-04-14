package ioKit

import (
	"github.com/richelieu42/chimera/v2/src/core/file/fileKit"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type (
	lumberjackOptions struct {
		maxSize    int
		maxAge     int
		maxBackups int
		localTime  bool
		compress   bool
		console    bool
	}

	LumberjackOption func(opts *lumberjackOptions)
)

func WithMaxSize(maxSize int) LumberjackOption {
	return func(opts *lumberjackOptions) {
		opts.maxSize = maxSize
	}
}

func WithMaxAge(maxAge int) LumberjackOption {
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

func WithConsole(console bool) LumberjackOption {
	return func(opts *lumberjackOptions) {
		opts.console = console
	}
}

func loadOptions(options ...LumberjackOption) *lumberjackOptions {
	opts := &lumberjackOptions{}
	for _, option := range options {
		option(opts)
	}
	return opts
}

func NewLumberjackWriteCloser(filePath string, options ...LumberjackOption) (io.WriteCloser, error) {
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}

	opts := loadOptions(options...)
	var writeCloser io.WriteCloser = &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    opts.maxSize,
		MaxBackups: opts.maxBackups,
		MaxAge:     opts.maxAge,
		LocalTime:  opts.localTime,
		Compress:   opts.compress,
	}
	if opts.console {
		writeCloser = MultiWriteCloser(writeCloser, NopCloserToWriter(os.Stdout))
	}
	return writeCloser, nil
}

// NewLumberjackWriteCloser1
/*
Deprecated: Use NewLumberjackWriteCloser instead.

PS:
(1) lumberjack: 日志切割组件.一般情况下，lumberjack配合其他日志库，实现日志的滚动(rolling)记录.
(2) 根据 maxBackups、maxAge 删除过期文件，根据 compress 决定是否压缩哪些未压缩的旧日志文件。

Golang 语言三方库 lumberjack 日志切割组件怎么使用？
	https://mp.weixin.qq.com/s/gGnovwzS1ucW3Afxcytp_Q
go语言的日志滚动(rolling)记录器——lumberjack
	https://zhuanlan.zhihu.com/p/430224518

@param filePath		日志路径，归档日志也会保存在对应目录下（若该值为空，则日志会保存到os.TempDir()目录下，日志文件名为<processname>-lumberjack.log）
					(1) 会尝试创建父级目录；
					(2) 文件不存在，会自动创建；
					(3) 文件存在，
						(3.1) 是文件，内容追加在最后；
						(3.2) 是目录，返回error.
@param maxSize		[单位: MB，默认: 100MB] 日志大小到达此值，就开始backup
@param maxBackups	旧日志保存的最大数量，默认保存所有旧日志文件
@param maxAge		[单位: days] 旧日志保存的最大天数，默认保存所有旧日志文件
@param localTime	[默认使用UTC时间] 是否使用本地时间戳？
@param compress		[默认: false] 对backup的日志是否进行压缩
*/
func NewLumberjackWriteCloser1(filePath string, maxSize, maxBackups, maxAge int, localTime, compress bool) (io.WriteCloser, error) {
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}

	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  localTime,
		Compress:   compress,
	}, nil
}
