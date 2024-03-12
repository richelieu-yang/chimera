package postgresqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/sql/gormKit"
	"gorm.io/driver/bigquery"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn e.g.	"bigquery://projectid/location/dataset"
				"bigquery://go-bigquery-driver/playground"
*/
func NewGormDB(dsn string, poolConfig *gormKit.PoolConfig, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := bigquery.Open(dsn)
	return gormKit.Open(dialector, poolConfig, opts...)
}
