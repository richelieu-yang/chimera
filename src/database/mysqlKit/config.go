package mysqlKit

import (
	"fmt"
	"gorm.io/gorm/logger"
	"time"
)

type (
	Config struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		// Host e.g."127.0.0.1:3306"
		Host   string `json:"host"`
		DBName string `json:"dbName"`
		//DsnConfig

		Log  LogConfig  `json:"log"`
		Pool PoolConfig `json:"pool"`
	}

	DsnConfig struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		// Host e.g."127.0.0.1:3306"
		Host   string `json:"host"`
		DBName string `json:"dbName"`
	}

	LogConfig struct {
		SlowThreshold time.Duration   `json:"slowThreshold"`
		LogLevel      logger.LogLevel `json:"logLevel"`
	}

	PoolConfig struct {
		MaxIdleConns    int           `json:"maxIdleConns"`
		MaxOpenConns    int           `json:"maxOpenConns"`
		ConnMaxLifetime time.Duration `json:"connMaxLifetime"`
	}
)

func (c *DsnConfig) ToDSN() string {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.UserName,
		c.Password,
		c.Host,
		c.DBName,
	)
}
