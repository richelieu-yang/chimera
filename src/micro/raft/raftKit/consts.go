package raftKit

import "time"

const (
	// DefaultApplyTimeout raft.Apply() 默认的超时时间.
	DefaultApplyTimeout = time.Second * 3
)
