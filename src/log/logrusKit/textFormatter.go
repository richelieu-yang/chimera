package logrusKit

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/core/timeKit"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

// NewTextFormatter
/*
PS: 外部在调用此方法后，建议调用: Logger.SetReportCaller(true)!!!

@param timestampFormat 可以为""（将采用默认值）
*/
func NewTextFormatter(timestampFormat string) logrus.Formatter {
	if strKit.IsEmpty(timestampFormat) {
		timestampFormat = string(timeKit.FormatEntire)
	}

	return &logrus.TextFormatter{
		/* 时间格式 */
		TimestampFormat: timestampFormat,
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

		CallerPrettyfier: func(f *runtime.Frame) (funcName string, fileName string) {
			s := strings.Split(f.Function, ".")
			funcName = s[len(s)-1]

			s1 := strKit.Split(f.File, "/")
			length := len(s1)
			if length >= 2 {
				fileName = fmt.Sprintf("%s/%s:%d", s1[length-2], s1[length-1], f.Line)
			} else {
				fileName = fmt.Sprintf("%s:%d", f.File, f.Line)
			}

			return funcName, fileName
		},
	}
}
