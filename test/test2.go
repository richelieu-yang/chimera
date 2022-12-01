package main

import (
	"fmt"
	"os"
)

func main() {
	//if err := os.Setenv("CCC", "1"); err != nil {
	//	panic(err)
	//}

	fmt.Println(os.LookupEnv("CCC"))
}
