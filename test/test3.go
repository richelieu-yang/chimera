package main

import "github.com/richelieu-yang/chimera/v2/src/cmdKit"

func main() {
	_ = cmdKit.NewCommand("java", []string{"-version"})

	//path := "lib/main.exe"
	//cmd := exec.Command("cmd.exe", "/c", "start", path)
	//
	//data, err := cmd.CombinedOutput()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Print(string(data))
}
