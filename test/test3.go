package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	if err := validateKit.Field("3.14", "numeric"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ok")
	}

	//if err := validateKit.Field(0, "required,min=-1,max=100"); err != nil {
	//	fmt.Println(err.Error()) // Key: '' Error:Field validation for '' failed on the 'required' tag
	//} else {
	//	fmt.Println("ok")
	//}
}
