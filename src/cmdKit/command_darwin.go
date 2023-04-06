package cmdKit

import (
	"os/exec"
	"syscall"
)

// NewCommand
/*
@param deathSig Mac环境不支持
*/
func NewCommand(setpgid bool, deathSig syscall.Signal, name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: setpgid,
		//Pdeathsig: deathSig,
	}
	return cmd
}
