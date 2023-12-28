package main

import (
	"bufio"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"io"
	"os"
	"time"
)

func main() {
	path := "/Users/richelieu/Documents/ino/notes.zip"
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	f = f

	start := time.Now()

	data, err := os.ReadFile(path)
	//data, err := readWithBuffer(f, int(dataSizeKit.MiB*64))
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(len(data))
	fmt.Println(dataSizeKit.ToReadableStringWithIEC(uint64(len(data))))
	fmt.Println(time.Since(start))
}

func readWithBuffer(f *os.File, bufSize int) ([]byte, error) {
	//reader := bufio.NewReader(f)
	reader := bufio.NewReaderSize(f, bufSize)
	buf := make([]byte, bufSize)

	var chunks []byte
	for {
		//从file读取到buf中
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		if n == 0 {
			break
		}

		//读取到最终的缓冲区中
		chunks = append(chunks, buf...)
	}
	return chunks, nil
}
