package fileKit

import (
	"github.com/zeromicro/go-zero/core/fs"
	"os"
)

// CloseOnExec makes sure closing the file on process forking.
/*
参考: go-zero中 fs.CloseOnExec.
*/
func CloseOnExec(f *os.File) {
	fs.CloseOnExec(f)
}
