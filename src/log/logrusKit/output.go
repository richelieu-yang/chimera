package logrusKit

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// defaultOutput 默认输出: 控制台
var defaultOutput io.Writer = os.Stdout

// SetOutput 修改logrus默认的输出
/*
PS: 但是 logrus.New() 返回的*logrus.Logger实例的Out属性仍旧是 os.Stderr，因此建议通过 NewLogger() 创建*logrus.Logger实例.
*/
func SetOutput(writer io.Writer) {
	if writer == nil {
		return
	}

	defaultOutput = writer
	logrus.SetOutput(defaultOutput)
}
