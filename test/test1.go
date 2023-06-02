package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type (
	DsnConfig struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
		// Host e.g."127.0.0.1:3306"
		Host   string `json:"host"`
		DBName string `json:"dbName"`
	}
)

func (c *DsnConfig) String() string {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.UserName,
		c.Password,
		c.Host,
		c.DBName,
	)
}

func main() {
	c := &DsnConfig{
		UserName: "root",
		Password: "~Test123",
		Host:     "127.0.0.1:3306",
		DBName:   "ccc2",
	}

	dsn := c.String()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var sqlDB *sql.DB
	sqlDB, err = db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	type User struct {
		Name     string
		Age      uint
		Birthday time.Time
	}
	user := &User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	db.Set("gorm:table_options", "ENGINE=MyISAM").AutoMigrate(&User{})

	// 自动建表
	if err := db.AutoMigrate(user); err != nil {
		panic(err)
	}
	// 插入数据
	result := db.Create(user) // 通过数据的指针来创建
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
