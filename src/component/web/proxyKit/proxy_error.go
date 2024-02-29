package proxyKit

import (
	"errors"
	"net"
)

// IsProxyDialError 代理请求返回的error，是否是因为dial目标地址失败？
func IsProxyDialError(err error) bool {
	opErr := &net.OpError{}
	if errors.As(err, &opErr) {
		return opErr.Op == "dial"
	}
	return false
}
