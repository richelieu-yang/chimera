package main

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/jsonKit"
)

func main() {
	//obj, err := jsonKit.Unmarshal([]byte(""), &JsonParams{})
	//fmt.Println(obj, err)
	//
	//obj, err = jsonKit.Unmarshal([]byte("{}"), JsonParams{})
	//fmt.Printf("%+v\n", err)
	//
	//obj, err = jsonKit.Unmarshal([]byte("1"), &JsonParams{})
	//fmt.Printf("%+v\n", err)

	jp := &JsonParams{}
	err := jsonKit.Unmarshal([]byte(""), jp)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

}
