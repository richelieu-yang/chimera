package mysqlKit

import (
	"database/sql"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model

	Name sql.NullString
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

	user := &User{
		Name: sql.NullString{
			String: "test",
			Valid:  true,
		},
	}
	if err := db.AutoMigrate(user); err != nil {
		logrus.Fatal(err)
	}

	rst := db.Where(User{
		Name: sql.NullString{
			String: "test",
			Valid:  true,
		},
	}).FirstOrCreate(user)
	if rst.Error != nil {
		panic(rst.Error)
	}
	// 创建记录，rst.RowsAffected == 1
	// 否则 rst.RowsAffected == 0
	fmt.Printf("RowsAffected: [%d]\n", rst.RowsAffected)
}
