package ioKit

//import (
//	"io"
//	"os"
//)
//
//func NewDualWriteCloser(wc, wc1 *io.WriteCloser) *multiWriteCloser {
//	io.MultiWriter(wc, wc1)
//}
//
//func NewFileWriterCloser(filePath string, toConsole bool) (io.WriteCloser, error) {
//
//	os.Stdout
//
//	f, err := os.Create(filePath)
//	if err != nil {
//		return nil, err
//	}
//
//	if !toConsole {
//		/* (1) 输出到：文件 */
//		return f, nil
//	}
//	/* (2) 输出到： 文件、控制台 */
//	return &multiWriteCloser{
//		f: f,
//		// 注意：此处两个传参的顺序不能颠倒
//		writer: io.MultiWriter(f, os.Stdout),
//	}, nil
//}
