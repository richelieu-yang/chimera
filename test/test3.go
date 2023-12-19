package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"time"
)

func main() {
	fmt.Println(validateKit.VarWithValue(time.Hour, time.Hour-time.Minute, "ltcsfield"))
	// Key: '' Error:Field validation for '' failed on the 'ltcsfield' tag
	fmt.Println(validateKit.VarWithValue(time.Hour, time.Hour+time.Minute, "ltcsfield"))
	// <nil>

	fmt.Println(validateKit.VarWithValue(time.Duration(0), -time.Minute, "omitempty,ltcsfield"))
	// <nil>
}
