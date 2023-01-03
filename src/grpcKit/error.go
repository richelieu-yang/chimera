package grpcKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IsDeadlineExceededError
/*
@param err 可以为nil
*/
func IsDeadlineExceededError(err error) bool {
	if err == nil {
		return false
	}

	// 防止多层error嵌套
	err = errorKit.Cause(err)

	if s, ok := status.FromError(err); ok {
		if s.Code() == codes.DeadlineExceeded {
			return true
		}
	}
	return false
}
