package fileKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"io"
	"os"
)

type (
	CustomizedFile struct {
		*os.File

		writeCloser io.WriteCloser
	}
)

func (rf *CustomizedFile) Write(b []byte) (int, error) {
	return rf.writeCloser.Write(b)
}

func (rf *CustomizedFile) WriteString(s string) (n int, err error) {
	return io.WriteString(rf.writeCloser, s)
}

func (rf *CustomizedFile) WriteAt(b []byte, off int64) (n int, err error) {
	if writerAt, ok := rf.writeCloser.(io.WriterAt); ok {
		return writerAt.WriteAt(b, off)
	}
	return 0, errorKit.New("not supported")
}

func (rf *CustomizedFile) Close() error {
	err := rf.writeCloser.Close()
	err1 := rf.File.Close()
	if err != nil {
		return err
	}
	return err1
}

// NewCustomizedFile 自定义的os.File（修改了输出）
/*
@param writeCloser 与 传参f 相关的io.WriteCloser实例
*/
func NewCustomizedFile(f *os.File, writeCloser io.WriteCloser) (*CustomizedFile, error) {
	//f, err := NewFileInAppendMode(filePath)
	//if err != nil {
	//	return nil, err
	//}

	if f == nil {
		return nil, errorKit.New("f == nil")
	}
	if writeCloser == nil {
		return nil, errorKit.New("writeCloser == nil")
	}
	return &CustomizedFile{
		File:        f,
		writeCloser: writeCloser,
	}, nil
}
