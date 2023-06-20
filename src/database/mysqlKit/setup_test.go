package mysqlKit

import (
	"database/sql"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"testing"
)

type User struct {
	ID int

	Name string `gorm:"not null"`
	Age  uint32
	Mail sql.NullString
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
	db.Set("gorm:table_options", "ENGINE=InnoDB") /*.AutoMigrate(&user{})*/

	user := &User{}
	if err := db.AutoMigrate(user); err != nil {
		logrus.Fatal(err)
	}

	rst := db.Create(&User{
		Name: "yjs",
		Age:  0,
		//Mail: "",
	})
	if rst.Error != nil {
		logrus.Fatal(rst.Error)
	}
	logrus.Infof("rst.RowsAffected: %d", rst.RowsAffected)

	//var s []User
	//rst := db.Find(&s)
	//if rst.Error != nil {
	//	logrus.Fatal(rst.Error)
	//}
	//logrus.Infof("%+v", s)
}
