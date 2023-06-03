//go:build !windows

package osKit

import (
	"github.com/richelieu42/chimera/v2/src/cmdKit"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/intKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

// GetMaxOpenFiles 同一时间最多可开启的文件数
/*
PS:
(1) 当前仅支持Mac、Linux环境.
(2) 为何使用 sh -c "ulimit -n" 而非 ulimit -n? https://www.thinbug.com/q/17483723
*/
func GetMaxOpenFiles() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ulimit -n")
	if err != nil {
		return 0, err
	}
	// e.g. "122880\n" => "122880"
	str = strKit.Trim(str)

	i, err := intKit.StringToInt(str)
	if err != nil {
		return 0, errorKit.Newf("result(%s) isn't a number", str)
	}
	return i, nil

	//cmd := exec.Command("sh", "-c", "ulimit -n")
	////cmd := exec.Command("ulimit", "-n")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	return 0, err
	//}
	//// strKit.Trim()是为了：去掉最后面的"\n"
	//str := strKit.Trim(out.String())
	//value, err := strconv.Atoi(str)
	//if err != nil {
	//	return 0, errorKit.New("result(%s) of command(%s) isn't a number", str, cmd.String())
	//}
	//return value, nil
}

// GetUserMaxProcesses 用户最多可开启的程序数目
/*
PS:
(1) 仅支持Mac、Linux环境；
(2) Process: 进程.
*/
func GetUserMaxProcesses() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ulimit -u")
	if err != nil {
		return 0, err
	}
	// e.g."5333\n" => "5333"
	str = strKit.Trim(str)

	i, err := intKit.StringToInt(str)
	if err != nil {
		return 0, errorKit.Newf("result(%s) isn't a number", str)
	}
	return i, nil

	//cmd := exec.Command("sh", "-c", "ulimit -u")
	////cmd := exec.Command("ulimit", "-u")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err := cmd.Run()
	//if err != nil {
	//	return 0, err
	//}
	//// strKit.Trim()是为了：去掉最后面的"\n"
	//str := strKit.Trim(out.String())
	//
	//value, err := strconv.Atoi(str)
	//if err != nil {
	//	return 0, errorKit.New("result(%s) of command(%s) isn't a number", str, cmd.String())
	//}
	//return value, nil
}
