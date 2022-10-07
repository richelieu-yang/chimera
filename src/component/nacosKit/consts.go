package nacosKit

import "github.com/richelieu42/go-scales/src/core/strKit"

type (
	NacosLogLevel *strKit.String
)

var (
	DebugLevel = NacosLogLevel(strKit.NewString("debug"))
	InfoLevel  = NacosLogLevel(strKit.NewString("info"))
	WarnLevel  = NacosLogLevel(strKit.NewString("warn"))
	ErrorLevel = NacosLogLevel(strKit.NewString("error"))
)
