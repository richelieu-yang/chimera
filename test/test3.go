package main

import (
	"encoding/base64"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
)

func main() {
	fmt.Println(base64Kit.DecodeStringToString("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9", base64Kit.WithEncoding(base64.RawURLEncoding)))
	fmt.Println(base64Kit.DecodeStringToString("eyJhIjoiYiIsImlhdCI6MTcwOTUzNTU5NX0", base64Kit.WithEncoding(base64.RawURLEncoding)))
}
