package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/office/excelKit"
)

func main() {
	f, err := excelKit.OpenFile("/Users/richelieu/Desktop/test.xls")
	if err != nil {
		panic(err) // unsupported workbook file format
	}
	fmt.Println(f)
}
