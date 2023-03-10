package goZeroKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
)

func NewDailyRotateRuleWriteCloser(filePath, delimiter string, days int, compress bool) (io.WriteCloser, error) {
	delimiter = strKit.EmptyToDefault(delimiter, "-")

	return logx.NewLogger(
		filePath,
		logx.DefaultRotateRule(
			filePath,
			delimiter,
			days,
			compress,
		),
		compress,
	)
}

// NewSizeLimitRotateRuleWriteCloser
/*
@param maxSize 单位: MB
*/
func NewSizeLimitRotateRuleWriteCloser(filePath, delimiter string, days, maxSize, maxBackups int, compress bool) (io.WriteCloser, error) {
	delimiter = strKit.EmptyToDefault(delimiter, "-")

	return logx.NewLogger(
		filePath,
		logx.NewSizeLimitRotateRule(
			filePath,
			filePath,
			days,
			maxSize,
			maxBackups,
			compress,
		),
		compress,
	)
}
