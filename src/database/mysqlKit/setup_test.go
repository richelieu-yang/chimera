package mysqlKit

import (
	"database/sql"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/jsonKit"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model

	Name sql.NullString
	Age  uint
}

func TestSetUp(t *testing.T) {
	if wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName); err != nil {
		panic(err)
	} else {
		fmt.Printf("wd: %s\n", wd)
	}

	type config struct {
		MySQL *Config `json:"mysql"`
	}

	c := &config{}
	confKit.MustLoad("chimera-lib/config.yaml", c)
	f, err := fileKit.NewFile("gorm.log")
	if err != nil {
		panic(err)
	}
	MustSetUp(c.MySQL, f)

	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	if err := db.AutoMigrate(&User{}); err != nil {
		logrus.Fatal(err)
	}

	//db.Create(&User{
	//	Name: sql.NullString{
	//		String: "test",
	//		Valid:  true,
	//	},
	//	Age: 100,
	//})

	var u User
	rst := db.First(&u)
	if rst.Error != nil {
		panic(rst.Error)
	}
	// 赋零值
	u.Age = 0

	fmt.Println(jsonKit.MarshalToString(u, jsonKit.WithIndent("    ")))

	m := mapKit.Encode(u)
	rst = db.Table("users").Updates(m)
	//rst = db.Model(&u).Updates(&u)
	if rst.Error != nil {
		panic(rst.Error)
	}

	//u.Name = sql.NullString{}
	//u.Age = 100
	//rst = db.Save(&u)
	//if rst.Error != nil {
	//	panic(rst.Error)
	//}
}
