package poolKit

import (
	"github.com/panjf2000/ants/v2"
)

// NewPool 创建 *ants.Pool 实例.
/*
PS:
(1) 通过 "func (p *Pool) Submit(task func()) error" 执行任务，任务无参无返回值.
(2) 默认情况下（即不传options）:
	(a) (Nonblocking: false) 阻塞模式
	(b) （MaxBlockingTasks: 0）阻塞模式下，最多因为调用Pool.Submit()而阻塞的goroutine数量 不设限制

@param options
*/
var NewPool func(size int, options ...ants.Option) (*ants.Pool, error) = ants.NewPool

// NewPoolWithFunc 创建 *ants.PoolWithFunc 实例.
/*
PS: 通过 "func (p *PoolWithFunc) Invoke(args interface{}) error" 执行任务，任务有1个参数（interface{}类型）无返回值.
*/
var NewPoolWithFunc func(size int, pf func(interface{}), options ...ants.Option) (*ants.PoolWithFunc, error) = ants.NewPoolWithFunc
