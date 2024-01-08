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
*/
func NewDefaultRaftNode(id, addr, dir string, fsm raft.FSM) (*raft.Raft, error) {
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
	config.Logger = hclog.New(&hclog.LoggerOptions{
		Name:   "raft",
		Level:  hclog.DefaultLevel,
		Output: hclog.DefaultOutput,
	})
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

	/* snapshotStore
	(1) raft.NewFileSnapshotStore(opts.dataDir, 1, os.Stderr)
	(2) raft.NewFileSnapshotStore(raftDir, 2, os.Stderr)
	retain: Control how many snapshots are retained. Must be at least 1.
	*/
	snapshotStore, err := raft.NewFileSnapshotStore(dir, 2, os.Stderr)
	if err != nil {
		return nil, err
	}

	/* transport
	(1) raft.NewTCPTransport(raftAddr, addr, 2, 5*time.Second, os.Stderr)
	(2) raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	*/
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	transport, err := raft.NewTCPTransport(addr, tcpAddr, 2, 5*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}

	return raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
}
