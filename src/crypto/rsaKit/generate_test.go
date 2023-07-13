package rsaKit

import (
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"testing"
)

func TestGenerateKeyFiles(t *testing.T) {
	_, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}

	bits := 2048
	format := PKCS1
	password := "dqwdqwd强无敌群多"

	priPath := "_pri.pem"
	pubPath := "_pub.pem"
	if err := GenerateKeyFiles(bits, format, password, priPath, pubPath); err != nil {
		panic(err)
	}
}
