package main

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/chai2010/webp"
)

func main() {
	var buf bytes.Buffer
	var data []byte
	var err error

	// 读取文件数据
	if data, err = ioutil.ReadFile("input.png"); err != nil {
		log.Println(err)
	}

	// 解码PNG图片
	img, err := webp.Decode(bytes.NewReader(data))
	if err != nil {
		log.Println(err)
	}

	// 编码为无损WebP
	if err = webp.Encode(&buf, img, &webp.Options{Lossless: true}); err != nil {
		log.Println(err)
	}

	// 写入到输出文件
	if err = ioutil.WriteFile("output.webp", buf.Bytes(), 0666); err != nil {
		log.Println(err)
	}
}
