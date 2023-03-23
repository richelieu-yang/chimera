package ocrKit

import (
	"github.com/richelieu42/chimera/src/cmdKit"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/pathKit"
	"github.com/richelieu42/chimera/src/core/sliceKit"
	"github.com/richelieu42/chimera/src/idKit"
)

//// GetText1 通过 第三方依赖otiai10/gosseract 实现
//// Deprecated Mac(M1)环境没跑通
///*
//!!!: 使用此函数，必须确保"CGO_ENABLED=1"，否则go run或go build会报错: undefined: gosseract.NewClient
//*/
//func GetText1(imgPath string, languages ...string) (string, error) {
//	if err := fileKit.AssertExistAndIsFile(imgPath); err != nil {
//		return "", err
//	}
//
//	client := gosseract.NewClient()
//	defer client.Close()
//	err := client.SetImage(imgPath)
//	if err != nil {
//		return "", err
//	}
//	err = client.SetLanguage(languages...)
//	if err != nil {
//		return "", err
//	}
//	return client.Text()
//}

// GetText 通过 命令行 实现
/*
PS:
(1) 实际上依赖于tesseract；
(2) 初步测试，此方法是并发安全的；
(3) 识别并不精确，有些图片可能识别不到文本（e.g.图片很小且就一两个字符）.
*/
func GetText(imgPath string, languages ...string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(imgPath); err != nil {
		return "", err
	}

	tempDir, err := pathKit.GetUniqueTempDir()
	if err != nil {
		return "", err
	}
	uuid := idKit.NewUUID()
	filePath := pathKit.Join(tempDir, uuid+".txt")
	defer fileKit.Delete(filePath)

	var lang string
	languages = sliceKit.RemoveEmpty(languages, true)
	if len(languages) == 0 {
		// 默认值
		lang = "chi_sim"
	} else {
		lang = sliceKit.Join(languages, "+")
	}

	result, err := cmdKit.ExecuteToString("tesseract", imgPath, pathKit.Join(tempDir, uuid), "-l", lang)
	if err != nil {
		return "", err
	}
	if !fileKit.Exist(filePath) {
		return "", errorKit.Simple(result)
	}
	if fileKit.IsDir(filePath) {
		return "", errorKit.Simple("filePath(%s) is a directory", filePath)
	}
	data, err := fileKit.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
