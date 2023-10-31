package tableKit

import "github.com/apcera/termtables"

// CreateTable 创建 *termtables.Table 实例.
var CreateTable func() *termtables.Table = termtables.CreateTable
