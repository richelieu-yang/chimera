package vipsKit

import (
	"testing"
)

func TestSetUp(t *testing.T) {
	SetUp(nil)

	if err := ToWebp("/Users/richelieu/Downloads/iShot_2023-10-16_11.07.52.png", "test.webp", nil); err != nil {
		panic(err)
	}
	if err := ToJpeg("/Users/richelieu/Downloads/iShot_2023-10-16_11.07.52.png", "test.jpg", nil); err != nil {
		panic(err)
	}
}
