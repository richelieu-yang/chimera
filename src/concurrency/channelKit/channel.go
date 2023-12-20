package channelKit

import "github.com/duke-git/lancet/v2/concurrency"

// NewChannel 返回一个Channel指针实例
func NewChannel[T any]() *concurrency.Channel[T] {
	return concurrency.NewChannel[T]()
}
