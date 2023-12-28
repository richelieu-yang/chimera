package fileKit

import (
	"io"
	"os"
)

// readFileWithSize
/*
PS:
(1) 代码修改于: os.ReadFile();
(2) Golang秒读32GB大文件，如何读取？ https://mp.weixin.qq.com/s/hx6F9ZXEh1g49apeuyiUBw
*/
func readFileWithSize(f *os.File, readSize int64) ([]byte, error) {
	if readSize < 1024 {
		// Richelieu: 至少要1KB 1KB地读.
		readSize = 1024
	}

	data := make([]byte, 0, readSize)
	for {
		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
		n, err := f.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}
	}
}
