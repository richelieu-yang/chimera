package postgresqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/sql/gormKit"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn e.g.
*/
func NewGormDB(dsn string, poolConfig *gormKit.PoolConfig, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := clickhouse.Open(dsn)
	return gormKit.Open(dialector, poolConfig, opts...)
}
