package zapKit

import (
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// StringToLevel
/*
@param defArgs 不传就采用默认值（INFO）
*/
func StringToLevel(str string, defArgs ...zapcore.Level) zapcore.Level {
	def := sliceKit.GetFirstItemWithDefault(zapcore.InfoLevel, defArgs...)

	switch strKit.ToLower(str) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return def
	}
}
