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

	//data, err := os.ReadFile(path)
	//data, err := ReadFile(path)
	//data, err := readWithBuffer(f, int(dataSizeKit.MiB*128))
	data, err := readWithBuffer(f, int(dataSizeKit.GiB))
	if err != nil {
		panic(err)
	}

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

// ReadFile reads the named file and returns the contents.
// A successful call returns err == nil, not err == EOF.
// Because ReadFile reads the whole file, it does not treat an EOF from Read
// as an error to be reported.
func ReadFile(name string) ([]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var size int
	if info, err := f.Stat(); err == nil {
		size64 := info.Size()
		if int64(int(size64)) == size64 {
			size = int(size64)
		}
	}
	size++ // one byte for final read at EOF

	// If a file claims a small size, read at least 512 bytes.
	// In particular, files in Linux's /proc claim size 0 but
	// then do not work right if read in small pieces,
	// so an initial read of 1 byte would not work correctly.
	if size < 512 {
		size = 512
	}

	//size := 1024

	data := make([]byte, 0, size)
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
