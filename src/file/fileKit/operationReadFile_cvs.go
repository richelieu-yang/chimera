package fileKit

import "github.com/duke-git/lancet/v2/fileutil"

func ReadCsvFile(filePath string) ([][]string, error) {
	if err := AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}

	return fileutil.ReadCsvFile(filePath)
}
