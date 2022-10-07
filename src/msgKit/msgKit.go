package msgKit

import (
	"github.com/richelieu42/go-scales/src/confKit"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
)

/*
key:	code
value:	msg
*/
var msgMap = make(map[string]string)

// ReadFile 读取message文件，加到 msgMap 中.
/*
@param filePath 建议是 .properties后缀 的文件
*/
func ReadFile(filePath string) error {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return err
	}

	m := make(map[string]string)
	if err := confKit.ReadFileAs(filePath, nil, &m); err != nil {
		return err
	}

	UpdateMsgMap(m)
	return nil
}

func Read(data []byte, fileType string) error {
	m := make(map[string]string)
	err := confKit.ReadAs(data, fileType, nil, &m)
	if err != nil {
		return err
	}

	UpdateMsgMap(m)
	return nil
}

func UpdateMsgMap(m map[string]string) {
	for code, msg := range m {
		msgMap[code] = msg
	}
}

func GetMsg(code string) string {
	return msgMap[code]
}
