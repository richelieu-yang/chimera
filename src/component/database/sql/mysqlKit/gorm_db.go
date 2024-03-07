package mysqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/sql/gormKit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormDB
/*
@param dsn 	参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
			想要正确的处理 time.Time ，您需要带上 parseTime 参数， (更多参数) 要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4
			e.g.
				"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
*/
func NewGormDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := mysql.Open(dsn)

	return gormKit.Open(dialector, opts...)
}
