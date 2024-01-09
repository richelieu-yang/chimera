package raftKit

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"net"
	"os"
	"path/filepath"
	"time"
)

// NewRaftNodeAndBootstrapCluster
/*
PS: 将 传参addr 作为id.

@param addr 	raft节点的地址（不能为""）
@param fsm 		不能为nil
@param logger 	节点的日志输出，可以为nil（将使用默认的logger，控制台 debug级别 由于默认配置）
*/
func NewRaftNodeAndBootstrapCluster(addr, dir string, fsm raft.FSM, logger hclog.Logger, nodeAddrs []string) (*raft.Raft, error) {
	if err := strKit.AssertNotEmpty(addr, "addr"); err != nil {
		return nil, err
	}
	if err := fileKit.AssertNotExistOrIsDir(dir); err != nil {
		return nil, err
	}
	if err := interfaceKit.AssertNotNil(fsm, "fsm"); err != nil {
		return nil, err
	}
	if err := validateKit.Var(nodeAddrs, "unique,gte=3,dive,hostname_port"); err != nil {
		return nil, errorKit.Wrap(err, "param nodeAddrs is invalid")
	}

	/* (0) config */
	// Richelieu: 此处不需要配置 config.NotifyCh，用不着
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(addr)
	config.Logger = logger
	config.SnapshotInterval = 20 * time.Second
	config.SnapshotThreshold = 2

	/* (1) logStore */
	logStore, err := raftboltdb.NewBoltStore(filepath.Join(dir, "raft-log.db"))
	if err != nil {
		return nil, err
	}

	/* (2) stableStore */
	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(dir, "raft-stable.db"))
	if err != nil {
		return nil, err
	}

	/* (3) snapshotStore */
	/*
		@param retain: Control how many snapshots are retained. Must be at least 1.

		e.g.
			raft.NewFileSnapshotStore(opts.dataDir, 1, os.Stderr)
		e.g.1
			raft.NewFileSnapshotStore(raftDir, 2, os.Stderr)
	*/
	snapshotStore, err := raft.NewFileSnapshotStore(dir, 2, os.Stderr)
	if err != nil {
		return nil, err
	}

	/* (4) transport */
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	/*
		@param maxPool		是一个整数，表示连接池的最大容量，用于复用TCP连接以减少开销。如果为0，则使用默认值2。
		@param timeout		是一个time.Duration类型，表示连接超时的时间，用于控制网络延迟的影响。如果为0，则使用默认值10 * time.Second。
		@param logOutput	是一个io.Writer接口，表示日志输出的目标，用于记录传输层的信息。如果为nil，则使用os.Stderr作为输出。

		e.g.
			raft.NewTCPTransport(raftAddr, addr, 2, 5*time.Second, os.Stderr)
		e.g.1
			raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	*/
	maxPool := 3
	timeout := 10 * time.Second
	transport, err := raft.NewTCPTransport(addr, tcpAddr, maxPool, timeout, os.Stderr)
	if err != nil {
		return nil, err
	}

	/* (5) raft node */
	node, err := raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return nil, err
	}

	/* (6) Bootstrap */
	var configuration raft.Configuration
	for _, addr := range nodeAddrs {
		server := raft.Server{
			ID:      raft.ServerID(addr),
			Address: raft.ServerAddress(addr),
		}
		configuration.Servers = append(configuration.Servers, server)
	}
	node.BootstrapCluster(configuration)

	return node, nil
}
