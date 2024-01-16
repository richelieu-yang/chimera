package gormKit

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// NewLogger
/*
参考: logger.Default
*/
func NewLogger(logConfig *LogConfig) logger.Interface {
	if logConfig == nil {
		logConfig = &LogConfig{
			Output:        os.Stdout,
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      0,
			Colorful:      true,
		}
	}

	writer := log.New(logConfig.Output, "\r\n", log.Ldate|log.Ltime|log.Lmicroseconds)
	return logger.New(writer, logger.Config{
		// 慢SQL阈值
		SlowThreshold: logConfig.SlowThreshold,
		// 日志级别
		LogLevel: logConfig.LogLevel,
		// 忽略 logger.ErrRecordNotFound（记录未找到错误） ？
		IgnoreRecordNotFoundError: false,
		// 彩色打印？
		Colorful: false,
	})
}

//func NewLogger1(writer logger.Writer, config logger.Config) logger.Interface {
//	return logger.New(writer, config)
//}
