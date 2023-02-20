package ioKit

import (
	"io"
	"os"
)

// CloseWriter 尝试关闭 传参writer（如果能关闭的话）
/*
PS: 如果 传参obj 实现了 io.Closer 接口，且非特殊值，尝试关闭它；否则直接返回nil.
*/
func CloseWriter(writer io.Writer) error {
	if closer, ok := writer.(io.Closer); ok {
		switch closer {
		case os.Stdout:
			fallthrough
		case os.Stderr:
			/* 2种特殊情况，不处理 */
			return nil
		default:
			/* 关闭 */
			return closer.Close()
		}
	}

	/* 未实现 io.Closer 接口，不处理，直接返回nil */
	return nil
}
