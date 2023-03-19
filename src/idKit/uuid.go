package idKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	uuid "github.com/satori/go.uuid"
)

// NewUUID uuid v4
/*
@return e.g. "fd794d5a-4e7d-456d-a9e4-e377bf00f0a0"
*/
func NewUUID() string {
	return uuid.NewV4().String()
}

// NewSimpleUUID uuid v4
/**
@return e.g. "e28351058d0c446b85e3d7896c87078b"
*/
func NewSimpleUUID() string {
	return strKit.ReplaceAll(NewUUID(), "-", "")
}
