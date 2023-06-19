package logrusKit

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// defaultOutput 默认输出: 控制台
var defaultOutput io.Writer = os.Stdout

// SetOutput 修改logrus默认的输出
func SetOutput(writer io.Writer) {
	if writer == nil {
		return
	}

	defaultOutput = writer
	logrus.SetOutput(defaultOutput)
}
