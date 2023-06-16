package mysqlKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

type user struct {
	Name     string
	Age      uint
	Birthday time.Time
}

func (u *user) TableName() string {
	return "user"
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

	user := &user{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	db.Set("gorm:table_options", "ENGINE=MyISAM") /*.AutoMigrate(&user{})*/
	// 自动建表
	if err := db.AutoMigrate(user); err != nil {
		logrus.Fatal(err)
	}
	// 插入数据
	result := db.Create(user) // 通过数据的指针来创建
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
