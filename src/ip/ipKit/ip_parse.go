package ipKit

import "net"

var (
	// ParseIPString string => net.IP
	ParseIPString func(s string) net.IP = net.ParseIP
)
