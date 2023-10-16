package vipsKit

import (
	"testing"
)

func TestSetUp(t *testing.T) {
	SetUp(nil)

	if err := ToHeif("", "test.heic", nil); err != nil {
		panic(err)
	}
}
