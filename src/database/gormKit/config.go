package gormKit

import (
	"time"
)

type (
	PoolConfig struct {
		MaxIdleConns    int           `json:"maxIdleConns"`
		MaxOpenConns    int           `json:"maxOpenConns"`
		ConnMaxLifetime time.Duration `json:"connMaxLifetime"`
	}
)
