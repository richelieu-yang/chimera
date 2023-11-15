package pushKit

import "github.com/richelieu-yang/chimera/v2/src/tableKit"

// GetStatistics
/*
e.g.
+----------+-------------------+
| Type     | Count of channels |
+----------+-------------------+
| idMap    | 0                 |
| bsidMap  | 0                 |
| userMap  | 0                 |
| groupMap | 0                 |
+----------+-------------------+
*/
func GetStatistics() string {
	table := tableKit.CreateTable()

	table.AddHeaders("Type", "Count of channels")

	table.AddRow("idMap", GetCountOfIdMap())
	table.AddRow("bsidMap", GetCountOfBsidMap())
	table.AddRow("userMap", GetCountOfUserMap())
	table.AddRow("groupMap", GetCountOfGroupMap())

	return table.Render()
}
