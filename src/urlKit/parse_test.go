package urlKit

import (
	"fmt"
	"net/url"
	"testing"
)

/*
go中url.ParseRequestURI和url.Parse函数的踩坑记		https://mp.weixin.qq.com/s/L8vNnzyyg6vMxXP1Hxii6Q
*/
func TestParse(t *testing.T) {
	rawURL := "http://localhost?wifi=true&carrier=#Staysafe AIS&os=android"

	u0, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	u1, err := url.ParseRequestURI(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(u0.String())
	fmt.Println(u1.String())
}

func TestPolyfillUrl(t *testing.T) {
	reqUrl := "http://127.0.0.1:8888/s/测试.wps?a=测试&b=b#中部"
	fmt.Println(reqUrl) // http://127.0.0.1:8888/s/测试.wps?a=测试&b=b#中部

	params := url.Values{
		"c": []string{"c"},
		"d": []string{"大"},
	}
	reqUrl1, err := PolyfillUrl(reqUrl, params)
	if err != nil {
		panic(err)
	}
	fmt.Println(reqUrl1) // http://127.0.0.1:8888/s/%E6%B5%8B%E8%AF%95.wps?a=%E6%B5%8B%E8%AF%95&b=b&c=c&d=%E5%A4%A7#%E4%B8%AD%E9%83%A8
}
