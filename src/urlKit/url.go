package urlKit

// PolyfillUrl 优化url（类似Chrome浏览器，地址栏中的url有问题，但复制出来的是好的）.
/*
@param queryParams 额外的query参数，	(1) 没有可传nil
									(2) 值中切片中的字符串应当是未处理（编码）过的
*/
var PolyfillUrl func(reqUrl string, queryParams map[string][]string) (string, error) = AddQueryParamsToUrl
