package nacosKit

import "github.com/richelieu42/chimera/src/core/ptrKit"

type (
	NacosLogLevel *string
)

var (
	DebugLevel = ptrKit.ToPtr("debug")
	InfoLevel  = ptrKit.ToPtr("info")
	WarnLevel  = ptrKit.ToPtr("warn")
	ErrorLevel = ptrKit.ToPtr("error")
)
