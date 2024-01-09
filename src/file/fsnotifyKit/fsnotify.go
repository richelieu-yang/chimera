package fsnotifyKit

import "github.com/fsnotify/fsnotify"

var (
	// NewWatcher 创建一个新的监控对象.
	/*
		PS: 不使用了要主动调用 Watcher.Close().
	*/
	NewWatcher func() (*fsnotify.Watcher, error) = fsnotify.NewWatcher
)
