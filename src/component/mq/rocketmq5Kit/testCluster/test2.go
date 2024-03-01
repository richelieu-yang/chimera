package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/crypto/caesarKit"
)

func main() {
	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.-BRTwjN-sAlUjO-82qDrNHdMtGAwgWH05PrN49Ep_sU"
	fmt.Println(str)
	fmt.Println(caesarKit.Encrypt(str, 10))
}
