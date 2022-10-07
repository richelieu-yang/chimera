package pathKit

import (
	"github.com/mitchellh/go-homedir"
	"github.com/richelieu42/go-scales/src/core/userKit"
)

// GetHomeDir 获取当前用户的home dir（当前用户主目录）
func GetHomeDir() string {
	return userKit.GetHomeDir()
}

// ExpandTilde "~"（波浪号） => 当前用户主目录
/*
e.g.
("")		=> "", nil
("~") 		=> "/Users/richelieu", nil
("~/a/b/c")	=> "/Users/richelieu/a/b/c", nil
*/
func ExpandTilde(path string) (string, error) {
	return homedir.Expand(path)
}
