package logrusKit

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// consoleOutput 控制台的输出
var consoleOutput io.Writer = os.Stdout

func SetConsoleWriter(writer io.Writer) {
	if writer == nil {
		return
	}
	consoleOutput = writer
	logrus.StandardLogger().Out = writer
}
