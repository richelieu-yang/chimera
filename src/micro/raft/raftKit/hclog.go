package raftKit

import "github.com/hashicorp/go-hclog"

// NewHcLogger
/*
PS: 返回值可以赋值给 raft.Config.Logger

@param opts 可以为nil（将采用默认值: os.Stderr、Info）
*/
var NewHcLogger func(opts *hclog.LoggerOptions) hclog.Logger = hclog.New
