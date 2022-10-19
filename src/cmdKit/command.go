package cmdKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"os/exec"
	"regexp"
)

// Execute 执行命令（会阻塞直到命令结束）
func Execute(name string, args ...string) ([]byte, error) {
	/*
		!!!: exec.Cmd结构体执行时，会处理路径中的空格（e.g. java可执行文件的绝对路径、-Djava.ext.dirs=的路径...）
	*/
	cmd := exec.Command(name, args...)

	return cmd.CombinedOutput()
}

// ExecuteToString 执行命令（会阻塞直到命令结束）
func ExecuteToString(name string, args ...string) (string, error) {
	data, err := Execute(name, args...)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// PolyfillCommandPath 处理命令行中的路径（主要是处理其中的空格）
/*
e.g.
("/root/a b/c  d/java") => `/root/a" "b/c"  "d/java`
*/
func PolyfillCommandPath(path string) string {
	re := regexp.MustCompile("\\s+")
	return re.ReplaceAllStringFunc(path, func(s string) string {
		return strKit.Format(`"%s"`, s)
	})
}
