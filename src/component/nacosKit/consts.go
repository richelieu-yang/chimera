package nacosKit

import "github.com/richelieu42/go-scales/src/core/strKit"

type (
	NacosLogLevel *string
)

var (
	DebugLevel = strKit.GetStringPtr("debug")
	InfoLevel  = strKit.GetStringPtr("info")
	WarnLevel  = strKit.GetStringPtr("warn")
	ErrorLevel = strKit.GetStringPtr("error")
)
