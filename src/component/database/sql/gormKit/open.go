package gormKit

import (
	"fmt"
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

	/* ping */
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// TODO: 要不要Close？
	defer sqlDB.Close()
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	if err := db.Raw("SELECT 1").Scan(&result); err != nil {
		return fmt.Errorf("ping database failed: %w", err)
	}

	return db, nil
}
