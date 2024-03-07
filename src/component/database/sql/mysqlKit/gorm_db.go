package mysqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/sql/gormKit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := mysql.Open(dsn)

	return gormKit.Open(dialector, opts...)
}
