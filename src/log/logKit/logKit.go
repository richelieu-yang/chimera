package logKit

import (
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"io"
	"log"
	"os"
)

// NewFileLogger
/*
@param logPath 如果文件不存在，会生成文件和目录；文件存在，新的内容会append
*/
func NewFileLogger(filePath, prefix string) (*log.Logger, error) {
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return nil, err
	}

	// 如果生成文件的话，其权限为0666（与os.Create()的权限一样）
	// os.OpenFile()的传参flag，可参考：https://studygolang.com/articles/22180
	writer, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	// 此处不能关闭writer，否则日志内容将写不进去
	//defer writer.CloseWriters()

	return newLogger(writer, prefix), nil
}

func newLogger(writer io.Writer, prefix string) *log.Logger {
	if strKit.IsNotEmpty(prefix) {
		// 为了把"prefix"和"日志正文"分隔开，加个空格
		prefix = strKit.AppendIfMissing(prefix, " ")
	}
	return log.New(writer, prefix, log.Ldate|log.Ltime|log.Lmicroseconds|log.Lmsgprefix)
}
