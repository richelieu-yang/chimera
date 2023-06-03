package main

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func main() {
	err := gerror.NewCode(gcode.New(10000, "", nil), "My Error")
	fmt.Println(err.Error())
	fmt.Println(gerror.Code(err))
}
