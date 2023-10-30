package fileKit

import "github.com/fsnotify/fsnotify"

// NewWatcher
/*
!!!: 第2个返回值为nil的情况下，第1个返回值: 如果不用了要手动关闭.

fsnotify（8.7k Star; 文件系统事件通知）
	https://github.com/fsnotify/fsnotify
*/
func NewWatcher(path string) (watcher *fsnotify.Watcher, err error) {
	if err = AssertExist(path); err != nil {
		return
	}

	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = watcher.Close()
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		return nil, err
	}
	return
}
