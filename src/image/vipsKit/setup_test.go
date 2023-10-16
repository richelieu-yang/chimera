package vipsKit

import (
	"testing"
)

func TestSetUp(t *testing.T) {
	SetUp(nil)

	if err := ToHeif("/Users/richelieu/Downloads/iShot_2023-10-16_11.07.52.png", "test.jpg", nil); err != nil {
		panic(err)
	}
}
