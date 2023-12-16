package mysqlKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"gorm.io/gorm"
)

// CreateDatabaseIfNotExists 自动创建数据库
/*
在 Go 项目中优雅的使用 gorm v2
	https://segmentfault.com/a/1190000039097157
*/
func CreateDatabaseIfNotExists(db *gorm.DB, database string) error {
	if err := interfaceKit.AssertNotNil(db, "db"); err != nil {
		return err
	}
	if err := strKit.AssertNotEmpty(database, "database"); err != nil {
		return err
	}

	// utf8mb4编码
	sqlStr := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4;", database)
	rst := db.Exec(sqlStr)
	return rst.Error
}
