package mysqlKit

import "gorm.io/gorm/logger"

type (
	writer struct {
	}
)

// NewLogger
/*
gorm Logger
	https://mp.weixin.qq.com/s/R4Q_nChfGBKpMk5WmTg9oQ
gorm 接管日志，使用自义定的输出方式
	https://www.likecs.com/show-307107493.html
GromV1 SQL输出到日志文件
	https://www.cnblogs.com/guanchaoguo/p/16448416.html
*/
func NewLogger() logger.Interface {
	return logger.Default

}
