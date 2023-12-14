package raftKit

import "github.com/hashicorp/raft"

var (
	// NewRaft 创建节点.
	/*
		@param conf 	*raft.Config		节点配置
		@param fsm 		raft.FSM			有限状态机（finite state machine）
		@param logs 	raft.LogStore		用来存储raft的日志
		@param stable 	raft.StableStore	稳定存储，用来存储raft集群的节点信息等
		@param snaps 	raft.SnapshotStore	快照存储，用来存储节点的快照信息
		@param trans 	raft.Transport		raft节点内部的通信通道
	*/
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
