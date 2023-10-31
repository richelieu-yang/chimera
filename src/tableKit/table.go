package tableKit

import "github.com/scylladb/termtables"

// CreateTable 创建 *termtables.Table 实例.
/*
(1) AddHeader()添加头部信息；
(2) AddRow()逐行添加数据；
(3) Render()返回渲染后的表格字符串.
*/
var CreateTable func() *termtables.Table = termtables.CreateTable
