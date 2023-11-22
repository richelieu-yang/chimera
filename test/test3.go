package main

import (
	"fmt"
	"strconv"
	"strings"
)

func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func main() {
	src := "hello 你好师姐~！@#￥%……&*（）——+「」：“《》？【】；‘，。、"

	dest := Encode(src)
	fmt.Println(dest)

	src1, err := Decode(dest)
	if err != nil {
		panic(err)
	}
	fmt.Println(src1)

	if src != src1 {
		panic("not equal")
	}
}

func Encode(str string) string {
	textQuoted := strconv.QuoteToASCII(str)
	// 去掉首尾的双引号
	return textQuoted[1 : len(textQuoted)-1]
}

func Decode(str string) (rst string, err error) {
	return strconv.Unquote(strings.Replace(strconv.Quote(str), `\\u`, `\u`, -1))
}
