// Package raftKit
/*
	Transport: raft集群内部节点之间的通信渠道，节点之间需要通过这个通道来进行日志同步、leader选举等.
	(1) TCPTransport: 		基于tcp，可以跨机器跨网络通信.
	(2) InmemTransport:		以内存为基础的传输层实现，它使用一个缓冲区来存存储请求和响应，并使用一个消费者通道来接收数据.
	(3) NetworkTransport:	综合了TCPTransport和InmemTransport的优点的传输层实现，它既支持TCP协议提供的可靠性和高效性，又支持内存提供的简单性和灵活性.
*/
package raftKit

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	"io"
	"net"
	"time"
)

var (
	NewTCPTransport func(bindAddr string, advertise net.Addr, maxPool int, timeout time.Duration, logOutput io.Writer) (*raft.NetworkTransport, error) = raft.NewTCPTransport

	NewTCPTransportWithConfig func(bindAddr string, advertise net.Addr, config *raft.NetworkTransportConfig) (*raft.NetworkTransport, error) = raft.NewTCPTransportWithConfig

	NewTCPTransportWithLogger func(bindAddr string, advertise net.Addr, maxPool int, timeout time.Duration, logger hclog.Logger) (*raft.NetworkTransport, error) = raft.NewTCPTransportWithLogger
)

var (
	NewInmemTransport            func(addr raft.ServerAddress) (raft.ServerAddress, *raft.InmemTransport)                        = raft.NewInmemTransport
	NewInmemTransportWithTimeout func(addr raft.ServerAddress, timeout time.Duration) (raft.ServerAddress, *raft.InmemTransport) = raft.NewInmemTransportWithTimeout
)

var (
	NewNetworkTransport           func(stream raft.StreamLayer, maxPool int, timeout time.Duration, logOutput io.Writer) *raft.NetworkTransport = raft.NewNetworkTransport
	NewNetworkTransportWithConfig func(config *raft.NetworkTransportConfig) *raft.NetworkTransport                                              = raft.NewNetworkTransportWithConfig
	NewNetworkTransportWithLogger func(stream raft.StreamLayer, maxPool int, timeout time.Duration, logger hclog.Logger) *raft.NetworkTransport = raft.NewNetworkTransportWithLogger
)
