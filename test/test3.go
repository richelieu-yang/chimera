package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
)

func main() {
	resp, err := http.Get("https://chat.bing.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 使用默认的探测器（DetectBest）尝试确定内容的编码
	charset, _, _ := charset.DetermineEncoding(body, "")
	fmt.Println(charset)

	//// 如果你知道原始数据的编码不是UTF-8，你可以创建一个转换器并读取解码后的文本
	//if !charset() {
	//	r := transform.NewReader(bytes.NewReader(body), charset.NewDecoder())
	//	body, err = ioutil.ReadAll(r)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	// 现在body是已知编码格式的字符串或字节切片，可以进一步处理了
}
