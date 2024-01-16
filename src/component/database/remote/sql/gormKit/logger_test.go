package gormKit

import (
	"context"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

func TestNewLogger(t *testing.T) {
	writer := log.New(os.Stdout, "\r\n", log.Ldate|log.Ltime|log.Lmicroseconds)

	// 参考: logger.Default
	l := logger.New(writer, logger.Config{
		// 慢SQL阈值
		SlowThreshold: 200 * time.Millisecond,
		// 日志级别
		LogLevel: logger.Info,
		// 忽略 logger.ErrRecordNotFound（记录未找到错误） ？
		IgnoreRecordNotFoundError: false,
		// 彩色打印？
		Colorful: true,
	})

	//l.Trace(context.TODO(), "")
	l.Info(context.TODO(), "info")
	l.Warn(context.TODO(), "warn")
	l.Error(context.TODO(), "error")
}
