package ioKit

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/pathKit"
	"github.com/richelieu42/chimera/src/core/timeKit"
	"io"
	"runtime"
	"time"
)

// NewRotateFileWriteCloser 超时根据: maxAge
/*
Deprecated: 推荐使用 NewLumberjackWriteCloser().

PS:
(0) 写是线程安全的；
(1) patternPath: 附带pattern的文件路径，e.g. "d:/test/test.%Y-%m-%d %H_%M_%S.log"
(2) 只会输出到文件，并不会输出到控制台；
(3) 第一个返回值，如果调用 CloseWriters() 后再调用 Write()，将返回error（invalid argument）.
(4) 如果filePath对应的文件已经存在，会追加在最后（并不会覆盖）.

@param softLinkFlag 	true: 生成软链接（替身）
@param toConsoleFlag 	true: 输出到文件日志的同时，也输出到控制台

e.g.
("aaa.log", time.Second*3, time.Second*30, true) => 最多同时存在 11 个日志文件（不算替身；30 / 3 + 1 = 11）
*/
func NewRotateFileWriteCloser(filePath string, rotationTime, maxAge time.Duration, softLinkFlag bool) (io.WriteCloser, error) {
	/* 默认值 */
	if rotationTime <= 0 {
		rotationTime = time.Hour * 12
	}
	if maxAge <= 0 {
		maxAge = timeKit.Week
	}

	options := []rotatelogs.Option{
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithMaxAge(maxAge),
	}
	if softLinkFlag {
		options = append(options, rotatelogs.WithLinkName(filePath))
	}

	return newRotateFileWriteCloser(filePath, options)
}

// NewRotateFileWriteCloserWithCount 超时根据: rotationCount
/*
Deprecated: 推荐使用 NewLumberjackWriteCloser().
*/
func NewRotateFileWriteCloserWithCount(filePath string, rotationTime time.Duration, rotationCount uint, softLinkFlag bool) (io.WriteCloser, error) {
	/* 默认值 */
	if rotationTime <= 0 {
		rotationTime = time.Hour * 12
	}
	if rotationCount <= 0 {
		rotationCount = 14
	}

	options := []rotatelogs.Option{
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(rotationCount),
	}
	if softLinkFlag {
		options = append(options, rotatelogs.WithLinkName(filePath))
	}

	return newRotateFileWriteCloser(filePath, options)
}

func newRotateFileWriteCloser(filePath string, options []rotatelogs.Option) (io.WriteCloser, error) {
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}

	pattern := getFilePattern(filePath)
	return rotatelogs.New(pattern, options...)
}

// getFilePattern 复用代码
/*
e.g. Mac M1
("/Users/richelieu/Downloads/111.log") => "111(2022-11-28T15-50-40).log"
*/
func getFilePattern(filePath string) string {
	dir := pathKit.GetParentDir(filePath)
	prefix := fileKit.GetPrefix(filePath)
	suffix := fileKit.GetSuffix(filePath)

	// Windows 和 Mac 的文件名不支持":"
	var timePattern string
	switch runtime.GOOS {
	case "windows":
		fallthrough
	case "darwin":
		// windows || darwin
		timePattern = "(%Y-%m-%dT%H-%M-%S)"
	case "linux":
		fallthrough
	default:
		// linux
		timePattern = "(%Y-%m-%dT%H:%M:%S)"
	}

	return pathKit.Join(dir, prefix+timePattern+suffix)
}
