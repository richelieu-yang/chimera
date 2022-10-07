package logrusKit

import (
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	"gitee.com/richelieu042/go-scales/src/core/timeKit"
	"github.com/sirupsen/logrus"
)

var DefaultTextFormatter logrus.Formatter

func init() {
	DefaultTextFormatter = NewTextFormatter(timeKit.EntireFormat)
}

func NewTextFormatter(timestampFormat timeKit.TimeFormat) logrus.Formatter {
	if strKit.IsEmpty(string(timestampFormat)) {
		return DefaultTextFormatter
	}
	return &logrus.TextFormatter{
		/* 时间格式 */
		TimestampFormat: string(timestampFormat),
		/* 禁止显示时间 */
		DisableTimestamp: false,
		/* 显示完整时间 */
		FullTimestamp: true,

		/* 禁止颜色显示 */
		DisableColors: true,
		ForceColors:   false,

		DisableQuote:     true,
		ForceQuote:       false,
		QuoteEmptyFields: false,
	}
}
