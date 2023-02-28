package urlKit

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"strconv"
)

type (
	Address struct {
		HostName string
		Port     int
	}
)

func (addr *Address) String() string {
	return fmt.Sprintf("%s:%d", addr.HostName, addr.Port)
}

func ParseToAddress(rawURL string) (*Address, error) {
	u, err := Parse(rawURL)
	if err != nil {
		return nil, err
	}

	hostName := u.Hostname()
	portStr := u.Port()
	if strKit.IsEmpty(portStr) {
		switch u.Scheme {
		case "http":
			portStr = "80"
			break
		case "https":
			portStr = "443"
			break
		default:
			return nil, errorKit.Simple("invalid rawURL(%s)", rawURL)
		}
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}
	return &Address{
		HostName: hostName,
		Port:     port,
	}, nil
}
