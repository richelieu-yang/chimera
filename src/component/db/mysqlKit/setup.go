package mysqlKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
)

var db *gorm.DB

func MustSetUp(config *Config, output io.Writer) {
	if err := SetUp(config, output); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// SetUp
/*
@param output 客户端的日志输出（nil: 输出到控制台）
*/
func SetUp(config *Config, output io.Writer) error {
	if config == nil {
		return errorKit.New("config == nil")
	}

	/* logger */
	if output == nil {
		output = os.Stdout
	}
	writer := log.New(output, "\r\n", log.Ldate|log.Ltime|log.Lmicroseconds)
	// 参考: logger.Default
	clientLogger := logger.New(writer, logger.Config{
		// 慢SQL阈值
		SlowThreshold: config.Log.SlowThreshold,
		// 日志级别
		LogLevel: config.Log.LogLevel,
		// 忽略 logger.ErrRecordNotFound（记录未找到错误） ？
		IgnoreRecordNotFoundError: false,
		// 彩色打印？
		Colorful: false,
	})

	tmpDB, err := gorm.Open(mysql.Open(config.ToDsnString()), &gorm.Config{
		Logger: clientLogger,
	})
	if err != nil {
		return err
	}

	/* verify */
	sqlDB, err := tmpDB.DB()
	if err != nil {
		return err
	}
	if err := sqlDB.Ping(); err != nil {
		_ = sqlDB.Close()
		return errorKit.Wrap(err, "fail to ping")
	}

	db = tmpDB
	return nil
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, errorKit.New("uninitialized component")
	}
	return db, nil
}
