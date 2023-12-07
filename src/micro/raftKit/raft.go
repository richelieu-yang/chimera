package raftKit

import "github.com/hashicorp/raft"

var NewRaft func(conf *raft.Config, fsm raft.FSM, logs raft.LogStore, stable raft.StableStore, snaps raft.SnapshotStore, trans raft.Transport) (*raft.Raft, error) = raft.NewRaft
