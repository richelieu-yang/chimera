package httpKit

const (
	// PlainContentType 纯文本格式.
	PlainContentType = "text/plain; charset=utf-8"

	// JsonContentType JSON数据格式.
	JsonContentType = "application/json; charset=utf-8"

	// XmlContentType XML数据格式.
	XmlContentType = "text/xml; charset=utf-8"

	// FormUrlencodedContentType 适用于 POST 请求.
	/*
		application/x-www-form-urlencoded: 将键值对的参数用&连接起来，如果有空格，将空格转换为+加号；有特殊符号，将特殊符号转换为ASCII HEX值.
		application/x-www-form-urlencoded是浏览器默认的编码格式，对于Get请求，是将参数转换?key=value&key=value格式，连接到url后.
	*/
	FormUrlencodedContentType = "application/x-www-form-urlencoded; charset=utf-8"

	// OctetStreamContentType 二进制流数据（如常见的文件下载）.
	/*
		参考：https://www.runoob.com/http/http-content-type.html
	*/
	OctetStreamContentType = "application/octet-stream; charset=utf-8"
)
