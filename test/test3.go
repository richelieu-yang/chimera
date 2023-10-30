package main

import (
	"bytes"
	"compress/flate"
	"fmt"
)

func main() {
	// 原始数据
	data := []byte("Hello, world!")

	// 创建一个缓冲区来保存压缩后的数据
	var buf bytes.Buffer

	// 创建一个新的flate.Writer
	w, err := flate.NewWriter(&buf, flate.DefaultCompression)
	if err != nil {
		panic(err)
	}

	// 将数据写入flate.Writer
	_, err = w.Write(data)
	if err != nil {
		panic(err)
	}

	// 关闭flate.Writer以确保所有压缩的数据都已写入缓冲区
	err = w.Close()
	if err != nil {
		panic(err)
	}

	// 打印压缩后的数据
	fmt.Println(buf.Bytes())
	fmt.Println(buf.String())
}
