package netKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
)

// ProcessAddresses
/*
@return 如果第二个返回值为nil，那么 len(第一个返回值) > 0
*/
func ProcessAddresses(addresses []string) ([]string, error) {
	if len(addresses) == 0 {
		return nil, errorKit.Simple("len(addresses) == 0")
	}
	addrs := sliceKit.Uniq(sliceKit.RemoveEmpty(addresses, true))
	if len(addrs) == 0 {
		return nil, errorKit.Simple("len(addrs) == 0")
	}

	for index, addr := range addrs {
		a, err := ParseToAddress(addr)
		if err != nil {
			return nil, err
		}
		addrs[index] = a.String()
	}
	return addrs, nil
}
