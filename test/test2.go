package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/cmdKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// GetUlimitInfo
/*
e.g. macOS
core file size          (blocks, -c) 0
data seg size           (kbytes, -d) unlimited
file size               (blocks, -f) unlimited
max locked memory       (kbytes, -l) unlimited
max memory size         (kbytes, -m) unlimited
open files                      (-n) 10240
pipe size            (512 bytes, -p) 1
stack size              (kbytes, -s) 8176
cpu time               (seconds, -t) unlimited
max user processes              (-u) 5333
virtual memory          (kbytes, -v) unlimited <nil>
*/
func GetUlimitInfo() (string, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ulimit -a")
	if err != nil {
		return "", err
	}
	str = strKit.TrimSpace(str)
	return str, nil
}

func main() {
	fmt.Println(GetUlimitInfo())
}
