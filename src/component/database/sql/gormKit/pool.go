package gormKit

import (
	"database/sql"
	"time"
)

type (
	PoolConfig struct {
		// MaxIdleConns <=0: no idle connections
		MaxIdleConns int `json:"maxIdleConns" yaml:"maxIdleConns"`

		// MaxOpenConns <= 0: there is no limit on the number of open connections
		MaxOpenConns int `json:"maxOpenConns" yaml:"maxOpenConns"`

		// ConnMaxLifetime <= 0: connections are not closed due to a connection's age
		ConnMaxLifetime time.Duration `json:"connMaxLifetime" yaml:"connMaxLifetime"`
	}
)

// TakeEffect 使连接池配置生效.
func (pc *PoolConfig) TakeEffect(sqlDB *sql.DB) {
	var c *PoolConfig
	if pc == nil {
		c = &PoolConfig{
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxLifetime: 30 * time.Minute,
		}
	} else {
		c = pc
	}

	// 设置空闲连接池中的最大连接数
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	// 设置数据库的最大打开连接数
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	// 设置连接的最大生命周期（超过此时间的连接将被关闭）
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(c.ConnMaxLifetime)
}
