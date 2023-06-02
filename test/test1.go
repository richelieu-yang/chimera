package main

import (
	"fmt"
	"github.com/klauspost/cpuid/v2"
)

func main() {
	fmt.Println(cpuid.CPU.VM())

}
