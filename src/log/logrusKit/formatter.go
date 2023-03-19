package logrusKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/core/timeKit"
	"github.com/sirupsen/logrus"
)

var DefaultTextFormatter = NewTextFormatter("")

// NewTextFormatter
/*
@param timeKit.FormatEntire 可以为""，将采用默认值
*/
func NewTextFormatter(timestampFormat timeKit.TimeFormat) logrus.Formatter {
	str := string(timestampFormat)
	if strKit.IsEmpty(str) {
		str = string(timeKit.FormatEntire)
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
