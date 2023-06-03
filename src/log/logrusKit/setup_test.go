package logrusKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	MustSetUp(&Config{
		Level:      "debug",
		PrintBasic: false,
	})

	//DisableQuoteTemporarily(nil, func() {
	//	logrus.Info("1\n2\n3\n")
	//})
	//logrus.Info("1\n2\n3\n")

	logrus.Info("---")
	logrus.Info(gfile.Pwd())
	logrus.Info(gfile.SelfDir())
	logrus.Info(gfile.SelfPath())
	logrus.Info(gfile.MainPkgPath())
	logrus.Info("---")

	//fmt.Println(gfile.Pwd())
	//fmt.Println(gfile.SelfDir())
	//fmt.Println(gfile.MainPkgPath())
}
