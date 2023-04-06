package fileKit

import (
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"os"
	"path/filepath"
)

// GetSize 获取文件（或目录）的大小
func GetSize(path string) (*dataSizeKit.DataSize, error) {
	if err := AssertExist(path); err != nil {
		return nil, err
	}

	if IsFile(path) {
		return GetFileSize(path)
	}
	return GetDirSize(path)
}

// GetFileSize 获取文件的大小
func GetFileSize(filePath string) (*dataSizeKit.DataSize, error) {
	if err := AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	size := &dataSizeKit.DataSize{
		Number: float64(info.Size()),
		Unit:   dataSizeKit.B,
	}
	return size.ToSuitableUint(), nil
}

// GetDirSize 获取目录的大小
/*
参考:
golang获取文件/目录（包含下面的文件）的大小: https://blog.csdn.net/n_fly/article/details/117080173
*/
func GetDirSize(dirPath string) (*dataSizeKit.DataSize, error) {
	if err := AssertExistAndIsDir(dirPath); err != nil {
		return nil, err
	}

	var bytes int64
	err := filepath.Walk(dirPath, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			bytes += info.Size()
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	size := &dataSizeKit.DataSize{
		Number: float64(bytes),
		Unit:   dataSizeKit.B,
	}
	return size.ToSuitableUint(), nil
}

//const (
//	// B 1
//	B SizeUnit = 1
//	// KB 1024
//	KB SizeUnit = 1 << 10
//	// MB 1048576 = 1024 * 1024
//	MB SizeUnit = 1 << 20
//	// GB 1073741824
//	GB SizeUnit = 1 << 30
//	// TB 1099511627776
//	TB SizeUnit = 1 << 40
//)
//
//type SizeUnit int
//
//func (unit SizeUnit) IntValue() int {
//	return int(unit)
//}

//// GetSizeByFilePath
///*
//！！！：不支持获取“目录”的大小，会返回0.
//*/
//func GetSizeByFilePath(filePath string, unit SizeUnit, places *intKit.Integer32) (float64, error) {
//	info, err := os.Stat(filePath)
//	if err != nil {
//		return 0, err
//	}
//
//	return GetSize(info.Size(), unit, places), nil
//}
//
//func GetSizeBySlice(s []byte, unit SizeUnit, places *intKit.Integer32) float64 {
//	return GetSize(int64(len(s)), unit, places)
//}
//
//// GetSize
///*
//@param length 文件内容([]byte)的长度
//@return 文件大小（单位：B || byte）
//*/
//func GetSize(length int64, unit SizeUnit, places *intKit.Integer32) float64 {
//	if places == nil {
//		return floatKit.Div(float64(length), float64(unit))
//	}
//	return floatKit.DivRound(places.GetValue(), float64(length), float64(unit))
//}

//// ToGB 文件大小单位转换， B => GB
//func ToGB(bytes float64, places *intKit.Integer32) float64 {
//	return floatKit.DivRound(places, bytes, float64(GB))
//}
//
//// ToMB 文件大小单位转换， B => MB
//func ToMB(bytes float64, places *intKit.Integer32) float64 {
//	return floatKit.DivWithScale(places, bytes, float64(MB))
//}
//
//// ToKB 文件大小单位转换， B => KB
//func ToKB(bytes float64, places *intKit.Integer32) float64 {
//	return floatKit.DivWithScale(places, bytes, float64(KB))
//}
