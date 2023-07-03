package main

import "github.com/richelieu-yang/chimera/v2/src/imageKit"

func main() {
	//s := bimg.IsImageTypeSupportedByVips(bimg.PDF)
	//fmt.Println(s.Load)
	//fmt.Println(s.Save)
	//
	//data, err := fileKit.ReadFile("/Users/richelieu/Desktop/a.pdf")
	//if err != nil {
	//	panic(err)
	//}
	//imgType := bimg.DetermineImageType(data)
	//fmt.Println(imgType)

	if err := imageKit.Convert("1.png", "1.pdf"); err != nil {
		panic(err)
	}
}
