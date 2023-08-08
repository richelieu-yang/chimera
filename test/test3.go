package main

import (
	"fmt"
	"github.com/klauspost/cpuid/v2"
	"strings"
)

func main() {
	fmt.Println("Name:", cpuid.CPU.BrandName)
	fmt.Println("PhysicalCores:", cpuid.CPU.PhysicalCores)
	fmt.Println("ThreadsPerCore:", cpuid.CPU.ThreadsPerCore)
	fmt.Println("LogicalCores:", cpuid.CPU.LogicalCores)
	fmt.Println("Family", cpuid.CPU.Family, "Model:", cpuid.CPU.Model, "Vendor ID:", cpuid.CPU.VendorID)
	fmt.Println("Features:", strings.Join(cpuid.CPU.FeatureSet(), ","))
	fmt.Println("Cacheline bytes:", cpuid.CPU.CacheLine)
	fmt.Println("L1 Data Cache:", cpuid.CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", cpuid.CPU.Cache.L1I, "bytes")
	fmt.Println("L2 Cache:", cpuid.CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", cpuid.CPU.Cache.L3, "bytes")
	fmt.Println("Frequency", cpuid.CPU.Hz, "hz")

	//// Test if we have these specific features:
	//if cpuid.CPU.Supports(SSE, SSE2) {
	//	fmt.Println("We have Streaming SIMD 2 Extensions")
	//}
}
