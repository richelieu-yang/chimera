package postgresqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/sql/gormKit"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn e.g."host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := postgres.Open(dsn)

	return gormKit.Open(dialector, opts...)
}
