package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"testing"
)

func TestNewLogger(t *testing.T) {
	/* 输出到控制台 */
	logger := NewLogger()
	logger.Info("to console")

	/* 输出到文件(rotatable) */
	path := "test.log"
	writer, err := ioKit.NewLumberjackWriteCloser(path)
	if err != nil {
		panic(err)
	}
	logger = NewLogger(WithWriter(writer))
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
