package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/urlKit"
)

func main() {
	reqUrl := "http://127.0.0.1:8888/s/测试.wps?a=测试&b=b#中部"
	//reqUrl := "http://127.0.0.1:8888/s/%E6%B5%8B%E8%AF%95.wps?a=%E6%B5%8B%E8%AF%95&b=b#%E4%B8%AD%E9%83%A8"

	fmt.Println(urlKit.PolyfillUrl(reqUrl))

	//u, err := url.Parse(reqUrl)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(u.String())
	//
	//fmt.Println(u.Query().Encode())
	//
	////// 编码后的 路由
	////fmt.Println("EscapedPath:\t", u.EscapedPath())
	////// 编码后的 Fragment
	////fmt.Println("EscapedFragment:\t", u.EscapedFragment())
	//////fmt.Println(u.ForceQuery)
	////fmt.Println(u.RawQuery)
	////fmt.Println(u.Query())
	//
	////fmt.Println(u.QueryEscape(reqUrl))
	////fmt.Println(u.PathEscape(reqUrl))
	//
	//{
	//	u, err := url.ParseRequestURI(reqUrl)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(u.String())
	//}
}

//func PolyfillRequestUrl(reqUrl string) (string, error) {
//	u, err := url.Parse(reqUrl)
//	if err != nil {
//		return "", err
//	}
//
//	rst := fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.EscapedPath())
//
//	// query
//	var queryText string
//	queryValues := u.Query()
//
//	// fragment
//	fragment := u.EscapedFragment()
//
//	mapKit.AssertNotEmpty()
//
//	return "", err
//}
