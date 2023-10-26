package wireKit

import "github.com/google/wire"

var (
	// NewSet 将多个提供者函数添加到一个集合中.
	/*
		PS:
		(1) 如果经常同时使用多个提供者函数，这非常有用;
		(2) 实践中， 一组业务相关的 provider 时常被放在一起组织成 ProviderSet，以方便维护与切换;
		(3) 可以将其他提供者函数集添加到提供者函数集中.
	*/
	NewSet func(...interface{}) wire.ProviderSet = wire.NewSet
)
