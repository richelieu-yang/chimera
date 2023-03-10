package ioKit

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/richelieu42/go-scales/src/consts"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/pathKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"io"
	"time"
)

// NewRotateFileWriteCloser rotationTime && maxAge
/*
Deprecated: 建立使用 writeCloser_RotateRule.go 里面的.

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
	/* 尝试创建父级目录 */
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

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
	return rotatelogs.New(toFilePathWithPattern(filePath), options...)
}

// NewRotateFileWriteCloser1 rotationTime && rotationCount
/*
Deprecated: 建立使用 writeCloser_RotateRule.go 里面的.
*/
func NewRotateFileWriteCloser1(filePath string, rotationTime time.Duration, rotationCount int, softLinkFlag bool) (io.WriteCloser, error) {
	/* 默认值 */
	if rotationTime <= 0 {
		rotationTime = time.Hour * 12
	}
	if rotationCount <= 0 {
		rotationCount = 14
	}

	options := []rotatelogs.Option{
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(uint(rotationCount)),
	}
	if softLinkFlag {
		options = append(options, rotatelogs.WithLinkName(filePath))
	}
	return rotatelogs.New(toFilePathWithPattern(filePath), options...)
}

// toFilePathWithPattern 复用代码
/*
e.g. Mac M1
("/Users/richelieu/Downloads/111.log") => "111(2022-11-28 15：50：40).log"
*/
func toFilePathWithPattern(filePath string) string {
	dir := pathKit.GetParentDir(filePath)
	prefix := fileKit.GetPrefix(filePath)
	suffix := fileKit.GetSuffix(filePath)

	// Windows 和 Mac 的文件名不支持":"
	timePattern := "(%Y-%m-%d %H" + consts.ColonInFileName + "%M" + consts.ColonInFileName + "%S)"
	return pathKit.Join(dir, prefix+timePattern+suffix)
}
