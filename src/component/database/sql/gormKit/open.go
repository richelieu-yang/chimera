package gormKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"gorm.io/gorm"
)

// Open 返回一个 *gorm.DB 实例.
func Open(dialector gorm.Dialector, opts ...gorm.Option) (*gorm.DB, error) {
	if err := interfaceKit.AssertNotNil(dialector, "dialector"); err != nil {
		return nil, err
	}
	if sliceKit.IsEmpty(opts) {
		opts = []gorm.Option{&gorm.Config{}}
	}

	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	// TODO: 要不要 Close()？
	//defer sqlDB.Close()

	return db, nil
}
