package msgKit

import (
	"github.com/richelieu42/chimera/v2/confKit"
	"github.com/richelieu42/chimera/v2/core/file/fileKit"
	"github.com/richelieu42/chimera/v2/core/mapKit"
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
	msgMap = mapKit.Merge(msgMap, m)
}

func GetMsg(code string) string {
	return msgMap[code]
}
