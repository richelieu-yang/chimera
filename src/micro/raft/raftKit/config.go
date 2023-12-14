package raftKit

import "github.com/hashicorp/raft"

var (
	// DefaultConfig 默认的配置.
	DefaultConfig func() *raft.Config = raft.DefaultConfig
)
