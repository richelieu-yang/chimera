## Windows环境，执行.exe文件
实际上是通过start命令.

golang 调用cmd执行EXE
    http://www.manongjc.com/detail/22-yljaaewtenbdcop.html
go 调用可执行程序并传参(windows 系统exe程序示例)
    https://blog.csdn.net/vily_lei/article/details/129985497

### demo（缺陷: lib/main.exe的输出丢了）
```golang
package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path := "lib/main.exe"
	cmd := exec.Command("cmd.exe", "/c", "start", path)

	data, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(data))
}
```
