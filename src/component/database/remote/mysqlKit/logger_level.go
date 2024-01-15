package mysqlKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"gorm.io/gorm/logger"
)

func StringToLevel(str string) logger.LogLevel {
	if strKit.IsEmpty(str) {

	}

	switch strKit.ToLower(str) {
	case "silent":

	case "error":

	case "warn":

	case "info":

	default:

	}
}
