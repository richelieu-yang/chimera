package mysqlKit

import (
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

	Name string
	Age  uint32
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

	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	db.Set("gorm:table_options", "ENGINE=MyISAM") /*.AutoMigrate(&user{})*/

	user := &User{
		Name: "Miro",
		Age:  18,
	}
	if err := db.AutoMigrate(user); err != nil {
		logrus.Fatal(err)
	}
	result := db.Omit("name").Create(user)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
