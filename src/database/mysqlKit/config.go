package mysqlKit

import (
	"fmt"
	"gorm.io/gorm/logger"
	"time"
)

type (
	Config struct {
		DsnConfig

		Log  LogConfig  `json:"log"`
		Pool PoolConfig `json:"pool"`
	}

	DsnConfig struct {
		UserName string `json:"userName"`
		Password string `json:"password,optional"`
		// Host e.g."127.0.0.1:3306"
		Host   string `json:"host"`
		DBName string `json:"dbName"`
	}

	LogConfig struct {
		SlowThreshold time.Duration   `json:"slowThreshold,default=200ms"`
		LogLevel      logger.LogLevel `json:"logLevel,default=4,options=1|2|3|4"`
	}

	PoolConfig struct {
		MaxIdleConns    int           `json:"maxIdleConns,default=32"`
		MaxOpenConns    int           `json:"maxOpenConns,default=128"`
		ConnMaxLifetime time.Duration `json:"connMaxLifetime,default=30m"`
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
