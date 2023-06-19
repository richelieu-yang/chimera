package logrusKit

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// output 控制台的输出
var output io.Writer = os.Stdout

func SetOutput(writer io.Writer) {
	if writer == nil {
		return
	}

	output = writer
	logrus.SetOutput(output)
}
