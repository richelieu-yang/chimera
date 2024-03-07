package postgresqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/sql/gormKit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormDB
/*
PS: TiDB 兼容 MySQL 协议.

@param dsn e.g."root:@tcp(127.0.0.1:4000)/test"
*/
func NewGormDB(dsn string, poolConfig *gormKit.PoolConfig, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := mysql.Open(dsn)
	return gormKit.Open(dialector, poolConfig, opts...)
}
