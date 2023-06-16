package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"testing"
)

func TestNewLogger(t *testing.T) {
	/* 输出到控制台 */
	logger := NewLogger()
	logger.Info("to consoleOutput")

	/* 输出到文件(rotatable) */
	path := "test.log"
	writer, err := ioKit.NewRotatableWriteCloser(path, 10*dataSizeKit.MiB)
	if err != nil {
		panic(err)
	}
	logger = NewLogger(WithOutput(writer))
	logger.Info("rotatable file logger")
}

func TestNewFileLogger(t *testing.T) {
	/* 输出到文件(not rotatable) */
	logger, err := NewFileLogger("test1.log")
	if err != nil {
		panic(err)
	}

	logger.Info("test")
	logger.Info("测试")
	logger.Info("~！@#￥%……&*（）——+`1234567890-=")
}
