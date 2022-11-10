package scalesLogKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

// newLogger
/*
@param prefix 每条输出语句的前缀，可以为""
*/
func newLogger(logrusLogger *logrus.Logger, output interface{}, filePath, prefix string) ILogger {
	if strKit.IsNotEmpty(prefix) && !strKit.EndWith(prefix, " ") {
		prefix = prefix + " "
	}

	switch output {
	case os.Stdout:
		filePath = "os.Stdout"
	case os.Stderr:
		filePath = "os.Stderr"
	default:
	}

	return &Logger{
		disposed: false,
		rwLock:   new(sync.RWMutex),

		prefix:       prefix,
		logrusLogger: logrusLogger,
		output:       output,
	}
}
