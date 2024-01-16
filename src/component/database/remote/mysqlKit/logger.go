package mysqlKit

import "gorm.io/gorm/logger"

func NewLogger(config *LogConfig) logger.Interface {

}

func NewLogger1(writer logger.Writer, config logger.Config) logger.Interface {
	return logger.New(writer, config)
}
