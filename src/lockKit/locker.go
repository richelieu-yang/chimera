package lockKit

import "github.com/gogf/gf/v2/os/gmlock"

// NewLocker 内存锁模块，支持按照给定键名动态生成互斥锁，并发安全并支持Try*Lock特性.
/*
参考: 内存锁-gmlock https://goframe.org/pages/viewpage.action?pageId=1114161

使用场景: 需要动态创建互斥锁，或者需要维护大量动态锁的场景.

PS:
(1) 当维护大量动态互斥锁的场景时，如果不再使用的互斥锁对象，请手动调用 Remove() 删除掉；
(2) 可以锁 "关键词"；
(3) 支持 "读写锁".
*/
func NewLocker() *gmlock.Locker {
	return gmlock.New()
}
