package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

func main() {
	var st sonyflake.Settings
	st.StartTime = time.Now()

	sf := sonyflake.NewSonyflake(st)

	for {
		fmt.Println(sf.NextID())
	}
}
