package poolKit

import (
	"github.com/panjf2000/ants/v2"
)

// NewPool 创建 *ants.Pool 实例.
/*
PS:
(1) 通过 Pool.Submit() 执行任务，任务无参无返回值.
(2) 默认情况下（即不传options）:
	(a) (Nonblocking: false) 阻塞模式
	(b) (MaxBlockingTasks: 0) 阻塞模式下，最多因为调用Pool.Submit()而阻塞的goroutine数量 不设限制.
(3) 可以通过 ants.WithLogger() 指定日志输出（默认输出到控制台）.

@param size (1) 如果<=0，生成的池是无限制的;
			(2) 即cap;
			(3) 并不是传了多少就立即创建多少协程.
*/
var NewPool func(size int, options ...ants.Option) (*ants.Pool, error) = ants.NewPool

// NewPoolWithFunc 创建 *ants.PoolWithFunc 实例.
/*
PS:
(1) 通过 Pool.Submit() 执行任务，任务无参无返回值.
(2) 默认情况下（即不传options）:
	(a) (Nonblocking: false) 阻塞模式
	(b) (MaxBlockingTasks: 0) 阻塞模式下，最多因为调用Pool.Submit()而阻塞的goroutine数量 不设限制.
(3) 可以通过 ants.WithLogger() 指定日志输出（默认输出到控制台）.

@param size (1) 如果<=0，生成的池是无限制的;
			(2) 即cap;
			(3) 并不是传了多少就立即创建多少协程.
*/
var NewPoolWithFunc func(size int, pf func(interface{}), options ...ants.Option) (*ants.PoolWithFunc, error) = ants.NewPoolWithFunc
