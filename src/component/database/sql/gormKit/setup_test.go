package gormKit

import (
	"database/sql"
	"gorm.io/gorm"
	"testing"
)

func TestSetUp(t *testing.T) {
	type User struct {
		gorm.Model

		Name sql.NullString
		Age  uint `gorm:"column:user_age"`
	}

	//if wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName); err != nil {
	//	panic(err)
	//} else {
	//	fmt.Printf("wd: %s\n", wd)
	//}
	//
	//type config struct {
	//	MySQL *Config `json:"mysql"`
	//}
	//
	//c := &config{}
	//confKit.MustLoad("_chimera-lib/config.yaml", c)
	//f, err := fileKit.Create("gorm.log")
	//if err != nil {
	//	panic(err)
	//}
	//MustSetUp(c.MySQL, f)
	//
	//db, err := GetDB()
	//if err != nil {
	//	panic(err)
	//}
	//// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	//db.Set("gorm:table_options", "ENGINE=InnoDB")
	//
	//if err := db.AutoMigrate(&User{}); err != nil {
	//	logrus.Fatal(err)
	//}
	//
	////db.Create(&User{
	////	Name: sql.NullString{
	////		String: "test",
	////		Valid:  true,
	////	},
	////	Age: 100,
	////})
	//
	//var u User
	//rst := db.First(&u)
	//if rst.Error != nil {
	//	panic(rst.Error)
	//}
	//
	////rst = db.Table("users").Where("id=?", 1).Updates(map[string]interface{}{
	////	"Age": 100,
	////})
	//rst = db.Model(&u).Updates(map[string]interface{}{
	//	"Age": 0,
	//})
	//if rst.Error != nil {
	//	panic(rst.Error)
	//}
}
