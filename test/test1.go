package main

import (
	"fmt"
	"github.com/klauspost/cpuid/v2"
)

func main() {
	// Apple
	fmt.Println(cpuid.CPU.VendorString)
	// Apple M1 Pro
	fmt.Println(cpuid.CPU.BrandName)
}
