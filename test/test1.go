package main

import (
	"encoding/base64"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

func main() {
	m := map[string]interface{}{
		"code":    0,
		"message": "success",
	}
	jsonData, err := jsonKit.Marshal(m)
	if err != nil {
		panic(err)
	}

	base64Str := base64Kit.EncodeToString(jsonData, base64Kit.WithEncoding(base64.StdEncoding))
	fmt.Println(base64Str) // eyJjb2RlIjowLCJtZXNzYWdlIjoic3VjY2VzcyJ9
}
