//go:build !windows

package cmdKit

import (
	"os/exec"
	"syscall"
)

// Setpgid
/*
参考:
PID, PPID, PGID与SID https://blog.csdn.net/Justdoit123_/article/details/101347971
*/
func Setpgid(cmd *exec.Cmd) {
	if cmd != nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
		}
	}
}
