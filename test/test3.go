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
