package msgKit

import (
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
)

/*
key:	code
value:	message
*/
var msgMap = make(map[string]string)

func ReadFiles(paths ...string) error {
	for _, path := range paths {
		if err := ReadFile(path); err != nil {
			return err
		}
	}
	return nil
}

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
	msgMap = mapKit.Merge(msgMap, m)
}

func GetMsg(code string) string {
	return msgMap[code]
}
