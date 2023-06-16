package main

import (
	"database/sql"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/database/mysqlKit"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func main() {
	c := &mysqlKit.DsnConfig{
		UserName: "root",
		Password: "~Test123",
		Host:     "127.0.0.1:3306",
		DBName:   "ccc2",
	}

	writer, err := NewGormLoggerWriter()
	if err != nil {
		logrus.Fatal(err)
	}
	db, err := gorm.Open(mysql.Open(c.ToDSN()), &gorm.Config{
		Logger: logger.New(writer, logger.Config{
			// 慢SQL阈值
			SlowThreshold: 200 * time.Millisecond,
			// 日志级别
			LogLevel: logger.Info,
			// 忽略 logger.ErrRecordNotFound（记录未找到错误） ？
			IgnoreRecordNotFoundError: false,
			// 彩色打印？
			Colorful: false,
		}),
	})
	if err != nil {
		logrus.Fatal(err)
	}

	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		logrus.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		logrus.Fatal(err)
	}

	type User struct {
		Name     string
		Age      uint
		Birthday time.Time
	}
	user := &User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	db.Set("gorm:table_options", "ENGINE=MyISAM") /*.AutoMigrate(&User{})*/

	// 自动建表
	if err := db.AutoMigrate(user); err != nil {
		logrus.Fatal(err)
	}
	// 插入数据
	result := db.Create(user) // 通过数据的指针来创建
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}

func NewGormLoggerWriter() (logger.Writer, error) {
	f, err := fileKit.NewFile("ccc.log")
	if err != nil {
		return nil, err
	}
	return log.New(f, "\r\n", log.Ldate|log.Ltime|log.Lmicroseconds), nil
}
