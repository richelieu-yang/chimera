package runtimeKit

import "github.com/sirupsen/logrus"

var manager func(hook func()) = logrus.RegisterExitHandler

func SetShutDownMonitor(p func(hook func())) bool {
	if p == nil {
		return false
	}

	manager = p
	return true
}

func AddShutDownHook(hook func()) {
	manager(hook)
}
