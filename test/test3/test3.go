package main

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/cmd/cmdKit"
	"os"
	"strconv"
)

func main() {
	var pid int
	args := os.Args[1:]
	if len(args) > 0 {
		var err error
		pid, err = strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
	} else {
		pid = 12345
	}
	fmt.Println("pid", pid)

	cmd := fmt.Sprintf(`top -n 1 -b -p %d  | grep %d | awk -F " " '{print $6}'`, pid, pid)
	fmt.Println("cmd", cmd)

	fmt.Println(cmdKit.ExecuteToString(context.TODO(), "sh", "-c", cmd))
}
