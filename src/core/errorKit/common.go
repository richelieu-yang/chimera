package errorKit

import (
	"github.com/richelieu42/go-scales/src/funcKit"
)

// addCallerInfoToFormat 在传参format前面加上: "$包名.$方法名"
func addCallerInfoToFormat(format string, extraSkips ...int) string {
	//extraSkip := sliceKit.GetFirstItemWithDefault(0, extraSkips...)

	var extraSkip int
	if extraSkips != nil {
		extraSkip = extraSkips[0]
	} else {
		extraSkip = 0
	}
	return funcKit.GetCallerNameWithSkip(3+extraSkip) + ": " + format
}
