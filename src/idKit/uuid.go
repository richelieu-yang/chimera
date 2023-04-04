package idKit

import (
	"github.com/gofrs/uuid/v5"
	"github.com/richelieu42/chimera/src/core/strKit"
)

// NewUUID uuid v4
/*
e.g.
() => 3064040f-b626-4d23-9d5d-de220e337d7a <nil>
*/
func NewUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// NewSimpleUUID uuid v4
/**
e.g.
() => a5140372dabd46f8bc3814b659d40708 <nil>
*/
func NewSimpleUUID() (string, error) {
	id, err := NewUUID()
	if err != nil {
		return "", err
	}
	id = strKit.ReplaceAll(id, "-", "")
	return id, nil
}
