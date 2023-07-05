package queueKit

import "github.com/gogf/gf/v2/container/gqueue"

// NewQueue 队列（先进先出）
/*
PS: 也支持固定队列大小，固定队列大小时队列效率和标准库的channel无异.

使用场景: 该队列是并发安全的，常用于多goroutine数据通信且支持动态队列大小的场景.
*/
var NewQueue func(limit ...int) *gqueue.Queue = gqueue.New
