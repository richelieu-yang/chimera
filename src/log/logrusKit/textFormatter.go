package logrusKit

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/core/timeKit"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

// NewTextFormatter
/*
PS: 外部在调用此方法后，建议调用: Logger.SetReportCaller(true)!!!

@param timestampFormat 可以为""（将采用默认值）

e.g. 日志输出
time=2023-03-23 16:46:23.398+08:00 level=info msg=[CHIMERA, PROCESS] pid: [8579]. func=PrintBasicDetails(logrusKit/basicDetails.go:17)
time=2023-03-23 16:46:23.398+08:00 level=info msg=[CHIMERA, OS] os: [darwin]. func=PrintBasicDetails(logrusKit/basicDetails.go:20)
time=2023-03-23 16:46:23.398+08:00 level=info msg=[CHIMERA, OS] arch: [arm64]. func=PrintBasicDetails(logrusKit/basicDetails.go:21)
*/
func NewTextFormatter(timestampFormat string) *logrus.TextFormatter {
	if strKit.IsEmpty(timestampFormat) {
		timestampFormat = string(timeKit.FormatEntire1)
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

		DisableQuote: false,
		ForceQuote:   true,

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

			// 把 file属性 整合到 func属性 里
			funcName = fmt.Sprintf("%s(%s)", funcName, fileName)
			// 不输出 file属性
			fileName = ""

			return funcName, fileName
		},
	}
}
