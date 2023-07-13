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

	options := []RsaOption{
		WithFormat(PKCS1),
		WithPassword(""),
	}
	priPath := "_pri.pem"
	pubPath := "_pub.pem"
	if err := GenerateKeyFiles(2048, priPath, pubPath, options...); err != nil {
		panic(err)
	}
}
