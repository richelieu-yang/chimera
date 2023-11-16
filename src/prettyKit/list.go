package prettyKit

import "github.com/jedib0t/go-pretty/v6/list"

// NewListWriter 格式化输出为 list 格式.
var NewListWriter func() list.Writer = list.NewWriter
