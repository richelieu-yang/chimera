/*
https://gorm.io/zh_CN/docs/connecting_to_the_database.html#MySQL
*/

package mysqlKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/database/sql/gormKit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewGormDB
/*
PS: 适用于 MySQL、MariaDB...

@param dsn 	参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
			想要正确的处理 time.Time ，您需要带上 parseTime 参数， (更多参数) 要支持完整的 UTF-8 编码，您需要将 charset=utf8 更改为 charset=utf8mb4
			e.g.
				"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
*/
func NewGormDB(dsn string, poolConfig *gormKit.PoolConfig, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := mysql.Open(dsn)
	return gormKit.Open(dialector, poolConfig, opts...)
}

// NewGormDBWithConfig MySQL 驱动程序提供了 一些高级配置 可以在初始化过程中使用.
/*
e.g. 传参mysqlConfig
mysql.Config{
  DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
  DefaultStringSize: 256, 				// string 类型字段的默认长度
  DisableDatetimePrecision: true, 		// 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
  DontSupportRenameIndex: true, 		// 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
  DontSupportRenameColumn: true, 		// 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
  SkipInitializeWithVersion: false, 	// 根据当前 MySQL 版本自动配置
}
*/
func NewGormDBWithConfig(mysqlConfig mysql.Config, poolConfig *gormKit.PoolConfig, opts ...gorm.Option) (*gorm.DB, error) {
	dialector := mysql.New(mysqlConfig)
	return gormKit.Open(dialector, poolConfig, opts...)
}
