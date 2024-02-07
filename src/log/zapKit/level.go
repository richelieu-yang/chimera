package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"go.uber.org/zap/zapcore"
)

func ParseLevel(str string) (zapcore.Level, error) {
	if strKit.IsBlank(str) {
		return zapcore.DebugLevel, nil
	}
	return zapcore.ParseLevel(str)

	//def := sliceKit.GetFirstItemWithDefault(zapcore.InfoLevel, defArgs...)
	//
	//switch strKit.ToLower(str) {
	//case "debug":
	//	return zap.DebugLevel
	//case "info":
	//	return zap.InfoLevel
	//case "warn":
	//	return zap.WarnLevel
	//case "error":
	//	return zap.ErrorLevel
	//default:
	//	return def
	//}
}
