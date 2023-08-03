package osKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
)

// GetThreadCount
/*
支持: 	Linux
不支持:	Mac
*/
func GetThreadCount() (int, error) {
	return 0, errorKit.New("not yet realized")

	//str, err := cmdKit.ExecuteToString("sh", "-c", "ps -eLf | wc -l")
	//if err != nil {
	//	return 0, err
	//}
	//str = strKit.TrimSpace(str)
	//
	//i, err := strconv.Atoi(str)
	//if err != nil {
	//	return 0, err
	//}
	//return i, nil
}
