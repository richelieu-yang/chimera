package raftKit

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	"io"
)

var (
	// NewDiscardSnapshotStore
	/*
		PS:
		(1) DiscardSnapshotStore: 不存储，忽略快照，相当于/dev/null（一般用于测试）
		(2) 实现了接口: raft.SnapshotStore

	*/
	NewDiscardSnapshotStore func() *raft.DiscardSnapshotStore = raft.NewDiscardSnapshotStore

	// NewFileSnapshotStore
	/*
		PS:
		(1) FileSnapshotStore: 文件持久化存储
		(2) 实现了接口: raft.SnapshotStore
	*/
	NewFileSnapshotStore           func(base string, retain int, logOutput io.Writer) (*raft.FileSnapshotStore, error) = raft.NewFileSnapshotStore
	NewFileSnapshotStoreWithLogger func(base string, retain int, logger hclog.Logger) (*raft.FileSnapshotStore, error) = raft.NewFileSnapshotStoreWithLogger

	// NewInmemSnapshotStore
	/*
		PS:
		(1) InmemSnapshotStore: 内存存储，不持久化，重启程序会丢失
		(2) 实现了接口: raft.SnapshotStore
	*/
	NewInmemSnapshotStore func() *raft.InmemSnapshotStore = raft.NewInmemSnapshotStore
)
