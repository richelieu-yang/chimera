package main

import (
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
)

func main() {
	_, err1 := gjson.Encode(func() {})
	err2 := gerror.Wrap(err1, `error occurred`)
	fmt.Printf("%+v", err2)
}
