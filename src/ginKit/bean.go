package ginKit

type (
	ResponsePackage struct {
		// HttpStatusCode 响应的http状态码（默认200，即 http.StatusOK）
		HttpStatusCode int

		// 响应文本
		Text string
		// 响应文件的路径
		FilePath string
		// 响应文件的内容
		FileContent []byte
		// 响应文件内容的类型（需要与 FileContent 搭配使用）
		ContentType string

		// FileName 文件名
		/*
			PS: 使用 FilePath 的情况下， FileName 是可选的，为空的话会自己从 FilePath 中截取
		*/
		FileName string
		Error    error
	}
)
