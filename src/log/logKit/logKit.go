package logKit

import (
	"gitee.com/richelieu042/go-scales/src/core/file/fileKit"
	"gitee.com/richelieu042/go-scales/src/core/file/rotateFileKit"
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	"io"
	"log"
	"os"
	"time"
)

// NewLogger
/*
@param logPath 如果文件不存在，会生成文件和目录；文件存在，新的内容会append
*/
func NewLogger(logPath, prefix string) (*log.Logger, error) {
	if err := fileKit.MkParentDirs(logPath); err != nil {
		return nil, err
	}

	// 如果生成文件的话，其权限为0666（与os.Create()的权限一样）
	// os.OpenFile()的传参flag，可参考：https://studygolang.com/articles/22180
	writer, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	// 此处不能关闭writer，否则日志内容将写不进去
	//defer writer.Close()

	return newLogger(writer, prefix), nil
}

func NewRotateLogger(logPath, prefix string, rotationTime, maxAge time.Duration) (*log.Logger, error) {
	writer, err := rotateFileKit.NewRotateWriter(logPath, rotationTime, maxAge)
	if err != nil {
		return nil, err
	}

	return newLogger(writer, prefix), nil
}

func newLogger(writer io.Writer, prefix string) *log.Logger {
	if strKit.IsNotEmpty(prefix) {
		// 为了把"prefix"和"日志正文"分隔开，加个空格
		prefix = strKit.AppendIfMissing(prefix, " ")
	}
	return log.New(writer, prefix, log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
}
