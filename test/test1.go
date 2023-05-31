package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAT time.Time `created_at`
}

func (User) TableName() string {
	return "users"
}

var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blogok?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = d
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func main() {
	var user User
	db.Debug().First(&user)
	fmt.Printf("user: %v\n", user)
}
