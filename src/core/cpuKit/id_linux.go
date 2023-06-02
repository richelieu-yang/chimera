package cpuKit

import (
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"os/exec"
)

// GetCpuId
/*
PS: cpu.Info() 目前仅支持Windows环境，不支持：Mac环境（M1）、Linux环境.

Linux下用命令查看CPU ID: https://blog.csdn.net/benwdm/article/details/84685292
*/
func GetCpuId() (string, error) {
	cmd := exec.Command("sh", "-c", "dmidecode -t 4 | grep ID |sort -u |awk -F': ' '{print $2}'")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	str := string(out)
	str = strKit.RemoveSpace(str)
	return str, nil
}
