package idKit

import (
	"github.com/google/uuid"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

// NewUUID UUIDv4
/*
@return 长度36

e.g.
	() => "936eff5f-97c6-4f8b-b26d-9bab1f65ff55"
*/
func NewUUID() string {
	return uuid.New().String()

	//id, err := uuid.NewRandom()
	//if err != nil {
	//	return "", err
	//}
	//return id.String(), nil
}

// NewSimpleUUID UUIDv4，去掉了其中所有"-"
/*
@return 长度32

e.g.
	() => "415ef754dc174b888b186873e093ced1"
*/
func NewSimpleUUID() string {
	return strKit.ReplaceAll(NewUUID(), "-", "")
}
