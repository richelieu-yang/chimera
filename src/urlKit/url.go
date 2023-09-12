package urlKit

import "net/url"

// Parse 解析url（http、https、rtsp、rtmp等协议）
/*
PS:
(1) URL.EscapedPath():		处理过的 path
(2) URL.EscapedFragment():	处理过的 fragment（# 后面的玩意）
(3) Values.Encode():		处理过的 query string
(4) url.Parse VS url.ParseRequestURI: 当要解析的url字符串中包含有字符“#”时，使用url.Parse解析，会导致#后面的参数解析不出来。而使用ParseRequestURI就能解析到.
	使用场景:
		(a) 不关心#后面的数据，使用: url.Parse
		(b) 关心#后面的数据，使用: url.ParseRequestURI

GoLand教程-Go URL解析
	https://mp.weixin.qq.com/s/i6uEUzvu5BPna5QtSYDSMQ
golang url.Parse 解析
	https://blog.csdn.net/zhuyuqiang1238/article/details/121807708

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
var Parse func(rawURL string) (*url.URL, error) = url.Parse

var ParseRequestURI func(rawURL string) (*url.URL, error) = url.ParseRequestURI

// PolyfillUrl 优化url（类似Chrome浏览器，地址栏中的url有问题，但复制出来的是好的）.
/*
@param queryParams 额外的query参数，	(1) 没有可传nil
									(2) 值中切片中的字符串应当是未处理（编码）过的
*/
func PolyfillUrl(reqUrl string, extraQueryParams map[string][]string) (string, error) {
	u, err := Parse(reqUrl)
	if err != nil {
		return "", err
	}

	/*
		!!!: 不要只使用 URL.String() ，原因: 该方法内部直接使用了 RawQuery 属性（满足条件的话），导致如果 RawQuery 中包含未处理字符（比如中文），返回值中还是会包含未处理字符
		TODO: 后续看官方是否会修改 URL.String() 中对query的处理.
	*/
	values := u.Query()
	AddToValues(values, extraQueryParams)
	u.RawQuery = values.Encode()
	return u.String(), nil
}
