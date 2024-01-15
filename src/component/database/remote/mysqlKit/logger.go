package mysqlKit

import "gorm.io/gorm/logger"

func NewLogger(writer logger.Writer, config logger.Config) logger.Interface {
	return logger.New(writer, config)
}
