package cmdKit

import (
	"os/exec"
	"syscall"
)

// NewCommand
/*
Golang SysProcAttr.Pdeathsig方法代码示例 https://vimsky.com/examples/detail/golang-ex-syscall-SysProcAttr-Pdeathsig-method.html

@param setpgid
@param deathSig
*/
func NewCommand(setpgid bool, deathSig syscall.Signal, name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   setpgid,
		Pdeathsig: deathSig,
	}
	return cmd
}
