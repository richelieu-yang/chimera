package gormKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"gorm.io/gorm/logger"
)

// StringToLevel
/*
4种日志级别:
	logger.Silent	屏蔽所有输出
	logger.Error	允许: error
	logger.Warn		允许: warn、error
	logger.Info		允许: info、warn、error
*/
func StringToLevel(str string) (level logger.LogLevel, err error) {
	switch strKit.ToLower(str) {
	case "silent":
		level = logger.Silent
	case "error":
		level = logger.Error
	case "warn":
		level = logger.Warn
	case "":
		fallthrough
	case "info":
		level = logger.Info
	default:
		err = errorKit.New("invalid str: %s", str)
	}
	return
}
