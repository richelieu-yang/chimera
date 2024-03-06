package gormKit

import (
	"fmt"
	"gorm.io/gorm/logger"
	"io"
	"time"
)

type (
	LogConfig struct {
		Output io.Writer

		SlowThreshold             time.Duration
		LogLevel                  logger.LogLevel
		Colorful                  bool
		IgnoreRecordNotFoundError bool
	}

	Config struct {
		DsnConfig

		Pool PoolConfig `json:"pool" yaml:"pool"`
	}

	DsnConfig struct {
		UserName string `json:"userName" yaml:"userName"`
		Password string `json:"password" yaml:"password"`
		// Host e.g."127.0.0.1:3306"
		Host   string `json:"host" yaml:"host"`
		DBName string `json:"dbName" yaml:"dbName"`
	}

	PoolConfig struct {
		MaxIdleConns    int           `json:"maxIdleConns" yaml:"maxIdleConns"`
		MaxOpenConns    int           `json:"maxOpenConns" yaml:"maxOpenConns"`
		ConnMaxLifetime time.Duration `json:"connMaxLifetime" yaml:"connMaxLifetime"`
	}
)

func (c *DsnConfig) GetDsnString() string {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.UserName,
		c.Password,
		c.Host,
		c.DBName,
	)
}
