package httpClientKit

//// UploadFile 上传单个文件
///*
//TODO: 待验证及测试，以及和上面的方法一起抽出通用代码.
//
//@param params 可以为nil
//*/
//func UploadFile(url string, params map[string]string, fileKey, filePath string) (statusCode int, data []byte, err error) {
//	if err = fileKit.ExistAndIsFile(filePath); err != nil {
//		return
//	}
//
//	// fileKey默认值: "file"
//	fileKey = strKit.EmptyToDefault(fileKey, "file")
//
//	// body
//	body := &bytes.Buffer{}
//	writer := multipart.NewWriter(body)
//	for k, v := range params {
//		// 此处无需对k、v进行编码处理
//		err = writer.WriteField(k, v)
//		if err != nil {
//			return
//		}
//	}
//	file, err := os.Open(filePath)
//	if err != nil {
//		return
//	}
//	defer file.Close()
//	formPart, err := writer.CreateFormFile(fileKey, fileKit.GetName(filePath))
//	if err != nil {
//		return
//	}
//	_, err = io.Copy(formPart, file)
//	if err != nil {
//		return
//	}
//	err = writer.Close()
//	if err != nil {
//		return
//	}
//
//	return post(url, body, writer.FormDataContentType())
//}
