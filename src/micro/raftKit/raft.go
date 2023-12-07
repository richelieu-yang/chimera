package raftKit

import "github.com/hashicorp/raft"

var (
	// NewRaft
	NewRaft func(conf *raft.Config, fsm raft.FSM, logs raft.LogStore, stable raft.StableStore, snaps raft.SnapshotStore, trans raft.Transport) (*raft.Raft, error) = raft.NewRaft

	// RecoverCluster
	RecoverCluster func(conf *raft.Config, fsm raft.FSM, logs raft.LogStore, stable raft.StableStore, snaps raft.SnapshotStore, trans raft.Transport, configuration raft.Configuration) error = raft.RecoverCluster

	// BootstrapCluster
	BootstrapCluster func(conf *raft.Config, logs raft.LogStore, stable raft.StableStore, snaps raft.SnapshotStore, trans raft.Transport, configuration raft.Configuration) error = raft.BootstrapCluster

	// GetConfiguration
	GetConfiguration func(conf *raft.Config, fsm raft.FSM, logs raft.LogStore, stable raft.StableStore, snaps raft.SnapshotStore, trans raft.Transport) (raft.Configuration, error) = raft.GetConfiguration

	// HasExistingState
	HasExistingState func(logs raft.LogStore, stable raft.StableStore, snaps raft.SnapshotStore) (bool, error) = raft.HasExistingState
)
