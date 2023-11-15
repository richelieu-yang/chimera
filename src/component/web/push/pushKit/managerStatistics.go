package pushKit

import "github.com/richelieu-yang/chimera/v2/src/tableKit"

func GetStatistics() string {
	table := tableKit.CreateTable()
	table.AddHeaders("Type", "Count of channels")

	table.AddRow("idMap", idMap.Size())
	table.AddRow("bsidMap", bsidMap.Size())
	userMap.RLockFunc(func() {
		var count int
		userMap.RLockFunc(func() {
			for _, userSet := range userMap.Map {
				userSet
			}
		})
		table.AddRow("userMap", count)
	})
	groupMap.RLockFunc(func() {
		var count int

		table.AddRow("groupMap", count)
	})

	return table.Render()
}
