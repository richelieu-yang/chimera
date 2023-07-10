package ioKit

import (
	"bufio"
	"io"
)

// CountLines 读取内容的行数
/*
PS:
(1) 如果最后一行只有'\n'，会被视作1行；
(2) 不要重复调用，否则第二次开始行数将不准确.
*/
func CountLines(r io.Reader) (count int, err error) {
	var (
		br        *bufio.Reader = bufio.NewReader(r)
		delimiter byte          = '\n'
	)

	for {
		_, err = br.ReadString(delimiter)
		if err != nil {
			if err == io.EOF {
				// 读取结束
				count++
				return
			}
			// 读取时出错
			count = 0
			return
		}
		// 读取还未结束
		count++
	}
}

// CountLines1 读取内容的行数
/*
PS:
(1) 如果最后一行只有'\n'，不会被视作1行；
(2) 不要重复调用，否则第二次开始行数将不准确.
*/
func CountLines1(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count, scanner.Err()
}
