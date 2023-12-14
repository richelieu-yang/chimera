package raftKit

import (
	"github.com/hashicorp/raft"
	"io"
	"net"
	"time"
)

var (
	// NewTCPTransport
	/*
		PS:
		(1) Transport是raft集群内部节点之间的通信渠道，节点之间需要通过这个通道来进行日志同步、leader选举等;
		(2) TCPTransport: 基于tcp，可以跨机器跨网络通信.
	*/
	NewTCPTransport func(bindAddr string, advertise net.Addr, maxPool int, timeout time.Duration, logOutput io.Writer) (*raft.NetworkTransport, error) = raft.NewTCPTransport
)
