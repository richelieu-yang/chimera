package memoryKit

import "runtime"

// GetMemoryStats 获取有关Go程序的实时内存统计信息.
/*
runtime.MemStats结构体的的字段:
(1) Alloc（单位: 字节）
	Alloc is bytes of allocated heap objects.
	分配的堆对象的字节数。

(2) TotalAlloc（单位: 字节）
	TotalAlloc is cumulative bytes allocated for heap objects.
	堆对象分配的累积字节数。与 Alloc 不同，当对象被释放时它不会减少。

(3) Sys（单位: 字节）
	Sys is the total bytes of memory obtained from the OS.
	从操作系统获得的内存总字节数。这测量了 Go 运行时为堆、栈和其他内部数据结构保留的虚拟地址空间。

(4) NumGC
	NumGC is the number of completed GC cycles.
	它表示自程序启动以来垃圾回收器运行的次数。每次垃圾回收器运行时，NumGC 的值都会增加。这个字段可以用来监控程序的垃圾回收情况。
	!!!: 如果你发现 NumGC 的值增长得非常快，那么可能意味着你的程序存在内存分配问题。

(5) EnableGC
	EnableGC indicates that GC is enabled. It is always true, even if GOGC=off.
	表示是否允许垃圾回收。它是一个 bool 类型，如果为 true，则允许垃圾回收；如果为 false，则禁止垃圾回收。

(6) DebugGC
	DebugGC is currently unused.
	表示是否启用调试垃圾回收。它是一个 bool 类型，如果为 true，则启用调试垃圾回收；如果为 false，则禁用调试垃圾回收。
*/
func GetMemoryStats() *runtime.MemStats {
	var stats = &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	return stats
}

func GetProgramMemoryStats() *runtime.MemStats {

}
