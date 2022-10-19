// Package rotateFileKit
/*
PS:
(1) 主要用于分割日志文件；
(2) 会自动创建父级目录；
(3) 返回的io.Write实例，可用于: log.New()、logrus.Logger 的Out属性.
*/
package rotateFileKit

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/pathKit"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"io"
	"time"
)

// NewRotateWriter
/*
@param patternPath 附带pattern的文件路径，e.g. "d:/test/test.%Y-%m-%d %H_%M_%S.log"
*/
func NewRotateWriter(filePath string, rotationTime, maxAge time.Duration, args ...bool) (io.Writer, error) {
	options := []rotatelogs.Option{
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithMaxAge(maxAge),
	}
	options = attachSoftLink(options, filePath, args...)

	return rotatelogs.New(toFilePathWithPattern(filePath), options...)
}

func NewRotateWriterWithCount(filePath string, rotationTime time.Duration, rotationCount uint, args ...bool) (io.Writer, error) {
	options := []rotatelogs.Option{
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(rotationCount),
	}
	options = attachSoftLink(options, filePath, args...)

	return rotatelogs.New(
		toFilePathWithPattern(filePath),
		options...,
	)
}

// toFilePathWithPattern
/*
e.g.
("d:/test999/test.log") => "d:\test999\test.%Y-%m-%d %H_%M_%S.log"
*/
func toFilePathWithPattern(filePath string) string {
	dir := pathKit.GetParentDir(filePath)
	prefix := fileKit.GetPrefix(filePath)
	suffix := fileKit.GetSuffix(filePath)
	return pathKit.Join(dir, prefix+".%Y-%m-%d %H_%M_%S"+suffix)
}

// attachSoftLink 由 args 决定是否生成软链接
func attachSoftLink(options []rotatelogs.Option, filePath string, args ...bool) []rotatelogs.Option {
	softLink := sliceKit.GetFirstItemWithDefault(false, args...)
	if softLink {
		options = append(options, rotatelogs.WithLinkName(filePath))
	}
	return options
}
