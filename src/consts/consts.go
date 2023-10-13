// Package consts 本库的全局常量.
package consts

import (
	"github.com/richelieu-yang/chimera/v2/internal/resources"
	"github.com/sirupsen/logrus"
)

const (
	// ProjectName 本库（项目）的名字
	ProjectName = "chimera"

	UpperProjectName = "CHIMERA"
)

var Banner string

func init() {
	data, err := resources.Asset("_resources/banner.txt")
	if err != nil {
		logrus.WithError(err).Fatalf("[%s] Fail to get banner.", UpperProjectName)
		return
	}
	Banner = string(data)
}
