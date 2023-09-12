package urlKit

import (
	"net/url"
)

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
var ParseQuery func(query string) (url.Values, error) = url.ParseQuery

// AddToValues
/*
@param queryParams !!!: 值中切片中的字符串应当是未处理（编码）过的
@return 必定不为nil
*/
func AddToValues(values url.Values, queryParams map[string][]string) url.Values {
	if values == nil {
		values = make(map[string][]string)
	}

	for k, s := range queryParams {
		for _, v := range s {
			values.Add(k, v)
		}
	}
	return values
}

// AddToRawQuery
/*
PS: 可能会修改 req.URL.RawQuery.
*/
func AddToRawQuery(u *url.URL, queryParams map[string][]string) {
	values := u.Query()
	values = AddToValues(values, queryParams)
	u.RawQuery = values.Encode()
}

//// ToQueryString Deprecated: use url.Values instead.
///*
//@param m 会对值进行 编码 操作
//@return 可能为""
//
//e.g.
//	(nil) => ""
//e.g.1
//	m := map[string]string{
//		"a": "test",
//		"b": "测试",
//	}
//	fmt.Println(urlKit.ToQueryString(m)) // a=test&b=%E6%B5%8B%E8%AF%95
//*/
//func ToQueryString(queryParams map[string]string) string {
//	var str string
//
//	for k, v := range queryParams {
//		// PS: k和v都有可能是""
//		if strKit.IsNotEmpty(str) {
//			str += "&"
//		}
//		if strKit.IsNotEmpty(v) {
//			str += k + "=" + EncodeURIComponent(v)
//		} else {
//			str += k
//		}
//	}
//	return str
//}
