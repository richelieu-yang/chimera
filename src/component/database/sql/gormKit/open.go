package gormKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"gorm.io/gorm"
)

// Open 返回一个 *gorm.DB 实例.
/*
@param options	(1) 可以为nil（将采用默认值）
				(2) 建议配置 gorm.Config.Logger，配置日志输出
*/
func Open(dialector gorm.Dialector, poolConfig *PoolConfig, options ...gorm.Option) (*gorm.DB, error) {
	if err := interfaceKit.AssertNotNil(dialector, "dialector"); err != nil {
		return nil, err
	}
	if sliceKit.IsEmpty(options) {
		options = []gorm.Option{&gorm.Config{}}
	}

	db, err := gorm.Open(dialector, options...)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// Richelieu: 此处不要Close，原因: sql.DB.Close() 源码的注释.
	//defer sqlDB.Close()

	/* ping */
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	/*
		连接池（GORM 使用 database/sql 来维护连接池） https://gorm.io/zh_CN/docs/connecting_to_the_database.html#%E8%BF%9E%E6%8E%A5%E6%B1%A0
	*/
	poolConfig.TakeEffect(sqlDB)

	return db, nil
}
