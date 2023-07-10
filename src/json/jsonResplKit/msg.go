package jsonResplKit

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

func readFiles(paths ...string) error {
	for _, path := range paths {
		if err := readFile(path); err != nil {
			return err
		}
	}
	return nil
}

// readFile 读取message文件，加到 msgMap 中.
/*
@param filePath 建议是 .properties 的文件
*/
func readFile(filePath string) error {
	if err := fileKit.AssertExistAndIsFile(filePath); err != nil {
		return err
	}

	m := make(map[string]string)
	if err := confKit.ReadFileAs(filePath, nil, &m); err != nil {
		return err
	}

	msgMap = mapKit.Merge(msgMap, m)
	return nil
}

func read(data []byte, fileType string) error {
	m := make(map[string]string)
	err := confKit.ReadAs(data, fileType, nil, &m)
	if err != nil {
		return err
	}

	msgMap = mapKit.Merge(msgMap, m)
	return nil
}
