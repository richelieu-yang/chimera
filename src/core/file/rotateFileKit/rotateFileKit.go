// Package rotateFileKit
/*
PS:
(1) 主要用于分割日志文件；
(2) 会自动创建父级目录；
(3) 返回的io.Write实例，可用于: log.New()、logrus.Logger 的Out属性.
*/
package rotateFileKit

import (
	"gitee.com/richelieu042/go-scales/src/core/file/fileKit"
	"gitee.com/richelieu042/go-scales/src/core/osKit"
	"gitee.com/richelieu042/go-scales/src/core/pathKit"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

// NewRotateWriter
/*
@param patternPath 附带pattern的文件路径，e.g. "d:/test/test.%Y-%m-%d %H_%M_%S.log"
*/
func NewRotateWriter(filePath string, rotationTime, maxAge time.Duration) (io.Writer, error) {
	options := []rotatelogs.Option{
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithMaxAge(maxAge),
	}
	// Linux环境下，才会创建软链接（防止在Goland中报错）
	if !osKit.IsWindows() {
		options = append(options, rotatelogs.WithLinkName(filePath))
	}

	return rotatelogs.New(
		toFilePathWithPattern(filePath),
		options...,
	)
}

func NewRotateWriterWithCount(filePath string, rotationTime time.Duration, rotationCount uint) (io.Writer, error) {
	options := []rotatelogs.Option{
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(rotationCount),
	}
	// Linux环境下，才会创建软链接
	if !osKit.IsWindows() {
		options = append(options, rotatelogs.WithLinkName(filePath))
	}

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
