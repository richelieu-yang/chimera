package goZeroKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
)

func NewDailyRotateRuleWriteCloser(filePath, delimiter string, days int, compress bool) (io.WriteCloser, error) {
	if days <= 0 {
		return nil, errorKit.Simple("invalid days(%d)", days)
	}
	delimiter = strKit.EmptyToDefault(delimiter, "-")

	rule := logx.DefaultRotateRule(
		filePath,
		delimiter,
		days,
		compress,
	)
	return logx.NewLogger(filePath, rule, compress)
}

// NewSizeLimitRotateRuleWriteCloser
/*
@param maxSize 单位: MB
*/
func NewSizeLimitRotateRuleWriteCloser(filePath, delimiter string, days, maxSize, maxBackups int, compress bool) (io.WriteCloser, error) {
	if days <= 0 {
		return nil, errorKit.Simple("invalid days(%d)", days)
	}
	if maxSize <= 0 {
		return nil, errorKit.Simple("invalid maxSize(%d)", maxSize)
	}
	if maxBackups <= 0 {
		return nil, errorKit.Simple("invalid maxBackups(%d)", maxBackups)
	}
	delimiter = strKit.EmptyToDefault(delimiter, "-")

	rule := logx.NewSizeLimitRotateRule(
		filePath,
		filePath,
		days,
		maxSize,
		maxBackups,
		compress,
	)
	return logx.NewLogger(filePath, rule, compress)
}
