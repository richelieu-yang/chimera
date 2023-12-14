package raftKit

import (
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
)

var (
	// NewBoltStore
	/*
		PS:
		(1) hashicorp提供了一个 raft-boltdb 来实现底层存储，它是一个嵌入式的数据库，能够持久化存储数据，我们直接用它来实现 LogStore 和 StableStore.
		(2) 实现了接口: raft.LogStore、raft.StableStore
	*/
	NewBoltStore func(path string) (*raftboltdb.BoltStore, error) = raftboltdb.NewBoltStore

	// NewInmemStore
	/*
		PS:
		(1) 实现了接口: raft.LogStore、raft.StableStore
	*/
	NewInmemStore func() *raft.InmemStore = raft.NewInmemStore
)
