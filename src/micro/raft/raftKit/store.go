package raftKit

import (
	"github.com/hashicorp/raft-boltdb"
)

var (
	NewBoltStore func(path string) (*raftboltdb.BoltStore, error) = raftboltdb.NewBoltStore
)
