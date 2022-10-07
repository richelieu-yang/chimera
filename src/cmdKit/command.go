package cmdKit

import (
	"os/exec"
)

// Execute 执行命令（会阻塞直到命令结束）
func Execute(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}

func ExecuteToString(name string, args ...string) (string, error) {
	data, err := Execute(name, args...)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
