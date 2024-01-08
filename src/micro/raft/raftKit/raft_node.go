package raftKit

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"net"
	"os"
	"path/filepath"
	"time"
)

// NewDefaultRaftNode
/*
@param id 		raft节点的id（可以为""，此时将使用 addr 作为 id）
@param addr 	raft节点的地址（不能为""）
@param fsm 		不能为nil
@param logger 	可以为nil（将使用默认的logger，debug级别 由于默认配置）
*/
func NewDefaultRaftNode(id, addr, dir string, fsm raft.FSM, logger hclog.Logger) (*raft.Raft, error) {
	if err := strKit.AssertNotEmpty(addr, "addr"); err != nil {
		return nil, err
	}
	id = strKit.EmptyToDefault(id, addr)
	if err := fileKit.AssertNotExistOrIsDir(dir); err != nil {
		return nil, err
	}
	if err := interfaceKit.AssertNotNil(fsm, "fsm"); err != nil {
		return nil, err
	}

	// Richelieu: 此处不需要配置 config.NotifyCh，用不着
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(id)
	config.Logger = logger
	config.SnapshotInterval = 20 * time.Second
	config.SnapshotThreshold = 2

	/* logStore */
	logStore, err := raftboltdb.NewBoltStore(filepath.Join(dir, "raft-log.db"))
	if err != nil {
		return nil, err
	}

	/* stableStore */
	stableStore, err := raftboltdb.NewBoltStore(filepath.Join(dir, "raft-stable.db"))
	if err != nil {
		return nil, err
	}

	/* snapshotStore */
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

	/* transport */
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

	return raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
}
