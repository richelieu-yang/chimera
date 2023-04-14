package ioKit

import (
	"io"
)

// TryToClose 尝试关闭（如果实现了io.Closer接口的话）.
/*
@param writers 	(1) 可以为nil（即不传参）
				(2) 其中可以有nil
@return 发生error的话，返回第一个
*/
func TryToClose(objs ...any) error {
	var err error

	for _, obj := range objs {
		if closer, ok := obj.(io.Closer); ok {
			tmp := closer.Close()
			if tmp != nil && err == nil {
				err = tmp
			}
		}
	}
	return err
}

// CloseCloser
/*
PS: 就算循环过程中返回了非nil的error，也要继续向下循环（尽可能多地关闭）.

@param closers (1) 可以为nil（即不传参）；(2) 其中可以有nil
@return 发生error的话，返回第一个
*/
func CloseCloser(closers ...io.Closer) error {
	var err error

	for _, closer := range closers {
		tmp := closer.Close()
		if tmp != nil && err == nil {
			err = tmp
		}
	}
	return err
}
