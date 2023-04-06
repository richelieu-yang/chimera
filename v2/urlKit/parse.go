package urlKit

import "net/url"

// Parse 解析url（http、https、rtsp、rtmp等协议）
/*
参考: https://blog.csdn.net/zhuyuqiang1238/article/details/121807708

@param rawURL !!!不能是如下格式: "localhost:8080"

e.g.
	u, err := urlKit.Parse("http://localhost:8080/go?a=123&b=456")
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)     // http
	fmt.Println(u.Host)       // localhost:8080
	fmt.Println(u.Hostname()) // localhost
	fmt.Println(u.Port())     // 8080
	fmt.Println(u.Path)       // /go
	fmt.Println(u.RawQuery)   // a=123&b=456
	fmt.Println(u.Query())    // map[a:[123] b:[456]]
*/
func Parse(rawURL string) (*url.URL, error) {
	return url.Parse(rawURL)
}

// ParseQuery
/*
e.g.
	u, err := urlKit.Parse("http://localhost:8080/go?a=123&b=456")
	if err != nil {
		panic(err)
	}
	fmt.Println(u.RawQuery) // a=123&b=456

	m, err := urlKit.ParseQuery(u.RawQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println(m) 			// map[a:[123] b:[456]]
*/
func ParseQuery(query string) (url.Values, error) {
	return url.ParseQuery(query)
}
