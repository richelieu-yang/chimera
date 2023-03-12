package fileKit

import (
	"os"
	"syscall"
)

// CloseOnExec makes sure closing the file on process forking.
/*
参考: go-zero中 fs.CloseOnExec.
*/
func CloseOnExec(f *os.File) {
	if f == nil {
		return
	}
	syscall.CloseOnExec(int(f.Fd()))
}
