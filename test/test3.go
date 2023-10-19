package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

type config struct {
	A int `json:"a,default=1"`
	B int `json:"b,default=2"`
	C int `json:"c"`
}

func main() {
	c := &config{}
	_, err := viperKit.Unmarshal([]byte("{}"), "json", nil, c)
	if err != nil {
		panic(err)
	}
	str, _ := jsonKit.MarshalIndentToString(c, "", "    ")
	fmt.Println(str)
}
