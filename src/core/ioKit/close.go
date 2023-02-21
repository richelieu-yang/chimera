package ioKit

import (
	"io"
)

// CloseWriters 尝试关闭 传参writer（如果能关闭的话）
/*
@param writers (1) 可以为nil（即不传参）；(2) 其中可以有nil
@return 发生error的话，返回第一个
*/
func CloseWriters(writers ...io.Writer) error {
	var err error

	for _, writer := range writers {
		if closer, ok := writer.(io.Closer); ok {
			tmp := closer.Close()
			if tmp != nil && err == nil {
				err = tmp
			}
		}
	}
	return err
}

// CloseReaders
/*
@param readers (1) 可以为nil（即不传参）；(2) 其中可以有nil
@return 发生error的话，返回第一个
*/
func CloseReaders(readers ...io.Reader) error {
	var err error

	for _, reader := range readers {
		if closer, ok := reader.(io.Closer); ok {
			tmp := closer.Close()
			if tmp != nil && err == nil {
				err = tmp
			}
		}
	}
	return err
}

// CloseClosers
/*
PS: 就算循环过程中返回了非nil的error，也要继续向下循环（尽可能多地关闭）.

@param closers (1) 可以为nil（即不传参）；(2) 其中可以有nil
@return 发生error的话，返回第一个
*/
func CloseClosers(closers ...io.Closer) error {
	var err error

	for _, closer := range closers {
		tmp := closer.Close()
		if tmp != nil && err == nil {
			err = tmp
		}
	}
	return err
}

func CloseWriteClosers(writeClosers ...io.WriteCloser) error {
	var err error

	for _, closer := range writeClosers {
		tmp := closer.Close()
		if tmp != nil && err == nil {
			err = tmp
		}
	}
	return err
}

func CloseReadClosers(readClosers ...io.ReadCloser) error {
	var err error

	for _, closer := range readClosers {
		tmp := closer.Close()
		if tmp != nil && err == nil {
			err = tmp
		}
	}
	return err
}
