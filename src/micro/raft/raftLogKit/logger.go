package raftLogKit

import "github.com/hashicorp/go-hclog"

// NewLogger
/*
PS: 返回值可以赋值给 raft.Config.Logger

@param opts 可以为nil（将采用默认值: os.Stderr、Info）
*/
var NewLogger func(opts *hclog.LoggerOptions) hclog.Logger = hclog.New
