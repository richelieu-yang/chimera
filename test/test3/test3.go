package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

func main() {
	headerValue := "fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5"
	langs := make([]string, 0)

	/*
		e.g.
		"zh-CN,zh;q=0.9"
		"zh,zh-TW;q=0.9,en-US;q=0.8,en;q=0.7,zh-CN;q=0.6"
		"fr-CH, fr;q=0.9, en;q=0.8, de;q=0.7, *;q=0.5"
	*/
	//headerValue := httpKit.GetHeader(req.Header, "Accept-Language")
	s := strKit.Split(headerValue, ",")
	for _, ele := range s {
		ele = strKit.TrimSpace(ele)
		ele = strKit.SubBeforeString(ele, ";q=")
		if strKit.IsEmpty(ele) {
			continue
		}
		langs = append(langs, ele)
	}
	fmt.Println(langs)
}
