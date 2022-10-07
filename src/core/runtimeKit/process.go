package runtimeKit

import (
	"github.com/shirou/gopsutil/v3/process"
	"os"
)

// PID 当前进程的id
var PID = os.Getpid()

// PidExists 判断 传参pid 对应的进程是否存在
/*
PS: 通过 第三方库gopsutil 实现.

e.g.
(-1) => 	(false, error(invalid pid -1))
(13120) =>	(true, nil)
(13121) =>	(false, nil)
*/
func PidExists(pid int32) (bool, error) {
	return process.PidExists(pid)
}
