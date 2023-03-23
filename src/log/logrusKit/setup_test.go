package logrusKit

import "testing"

func TestSetUp(t *testing.T) {
	SetUp(&Config{
		Level:      "",
		PrintBasic: true,
	})

}
