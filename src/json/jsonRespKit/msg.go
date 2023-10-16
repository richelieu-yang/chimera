package jsonRespKit

import (
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
)

type FileData struct {
	Data []byte
	// FileType 文件的类型，e.g. "properties"、"yaml"、"ini"...
	FileType string
}

/*
key:	code
value:	message
*/
var msgMap = make(map[string]string)

// readFile 读取message文件，加到 msgMap 中.
/*
@param filePath 建议是 .properties 的文件
*/
func readFile(filePath string) error {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return err
	}

	m := make(map[string]string)
	if err := viperKit.ReadFileAs(filePath, nil, &m); err != nil {
		return err
	}
	msgMap = mapKit.Merge(msgMap, m)
	return nil
}

func readFileData(fd *FileData) error {
	if fd == nil {
		return nil
	}

	m := make(map[string]string)
	if err := viperKit.ReadAs(fd.Data, fd.FileType, nil, &m); err != nil {
		return err
	}
	msgMap = mapKit.Merge(msgMap, m)
	return nil
}
