package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	// GetModificationTime 获取文件（或目录）的修改时间
	GetModificationTime = gfile.MTime

	// GetModificationTimestamp 获取文件（或目录）的修改时间（单位: s）
	GetModificationTimestamp = gfile.MTimestamp

	// GetModificationTimestampMilli 获取文件（或目录）的修改时间（单位: ms）
	GetModificationTimestampMilli = gfile.MTimestampMilli
)

//// GetModificationTime 获取文件（或目录）的修改时间
///*
//@param path 传参""将返回err（Stat : The system cannot find the path specified.）
//*/
//func GetModificationTime(path string) (time.Time, error) {
//	info, err := Stat(path)
//	if err != nil {
//		return time.Time{}, err
//	}
//	return info.ModTime(), nil
//}
