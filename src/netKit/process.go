package netKit

import (
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

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
