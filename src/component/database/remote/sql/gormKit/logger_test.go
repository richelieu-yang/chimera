package gormKit

import (
	"context"
	"gorm.io/gorm/logger"
	"os"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	l := NewLogger(&LogConfig{
		Output:                    os.Stdout,
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		Colorful:                  true,
		IgnoreRecordNotFoundError: false,
	})

	//l.Trace(context.Background(),)
	l.Info(context.TODO(), "info")
	l.Warn(context.TODO(), "warn")
	l.Error(context.TODO(), "error")
}
