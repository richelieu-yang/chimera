package httpKit

const (
	// PlainContentType 纯文本格式
	PlainContentType = "text/plain; charset=utf-8"

	// JsonContentType JSON数据格式
	JsonContentType = "application/json; charset=utf-8"

	// XmlContentType XML数据格式
	XmlContentType = "text/xml; charset=utf-8"

	// OctetStreamContentType 二进制流数据（如常见的文件下载）
	/*
		参考：https://www.runoob.com/http/http-content-type.html
	*/
	OctetStreamContentType = "application/octet-stream; charset=utf-8"
)
