package ioKit

import (
	"io"
	"os"
)

type (
	// fileConsoleWriteCloser 输出到： 文件、控制台
	fileConsoleWriteCloser struct {
		f      *os.File
		writer io.Writer
	}
)

func (wc *fileConsoleWriteCloser) Write(p []byte) (int, error) {
	return wc.writer.Write(p)
}

func (wc *fileConsoleWriteCloser) Close() error {
	return wc.f.Close()
}

func NewFileWriterCloser(filePath string, toConsole bool) (io.WriteCloser, error) {
	f, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	if !toConsole {
		/* (1) 输出到：文件 */
		return f, nil
	}
	/* (2) 输出到： 文件、控制台 */
	return &fileConsoleWriteCloser{
		f: f,
		// 注意：此处两个传参的顺序不能颠倒
		writer: io.MultiWriter(f, os.Stdout),
	}, nil
}
